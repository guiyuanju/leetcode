package main

import "fmt"

func main() {
	root1 := makeBinaryTree([]any{3, 5, 1, 6, 2, 9, 8, nil, nil, 7, 4})
	root2 := makeBinaryTree([]any{3, 5, 1, 6, 7, 4, 2, nil, nil, nil, nil, nil, nil, 9, 8})
	assertEq(leafSimilar(root1, root2), true)

	root1 = makeBinaryTree([]any{1, 2, 3})
	root2 = makeBinaryTree([]any{1, 3, 2})
	assertEq(leafSimilar(root1, root2), false)
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	left := leaf(root1)
	right := leaf(root2)
	if len(left) != len(right) {
		return false
	}
	for i := range left {
		if left[i] != right[i] {
			return false
		}
	}
	return true
}

func leaf(root *TreeNode) []int {
	var res []int
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Left == nil && root.Right == nil {
			res = append(res, root.Val)
			return
		}
		if root.Left != nil {
			dfs(root.Left)
		}
		if root.Right != nil {
			dfs(root.Right)
		}
	}

	dfs(root)

	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func makeBinaryTree(nodes []any) *TreeNode {
	var root *TreeNode
	queue := []*TreeNode{}
	for i, n := range nodes {
		var cur *TreeNode
		if n != nil {
			cur = &TreeNode{
				Val:   n.(int),
				Left:  nil,
				Right: nil,
			}
		}
		queue = append(queue, cur)

		for queue[0] == nil {
			queue = queue[1:]
		}

		if i == 0 {
			root = cur
		} else if i%2 != 0 {
			queue[0].Left = cur
		} else {
			queue[0].Right = cur
			queue = queue[1:]
		}
	}
	return root
}

func printBinaryTree(root *TreeNode) {
	flat := []*TreeNode{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		flat = append(flat, cur)
		if cur != nil {
			queue = append(queue, cur.Left)
			queue = append(queue, cur.Right)
		}
	}
	for i := len(flat) - 1; i >= 0 && flat[i] == nil; i-- {
		flat = flat[:len(flat)-1]
	}
	fmt.Print("[")
	for i, v := range flat {
		if v != nil {
			fmt.Print(v.Val)
		} else {
			fmt.Print("null")
		}
		if i < len(flat)-1 {
			fmt.Print(", ")
		}
	}
	fmt.Println("]")
}

func assertEq[T comparable](a, b T) {
	if a != b {
		fmt.Printf("Failed: %v != %v\n", a, b)
	} else {
		fmt.Printf("Ok: %v\n", a)
	}
}
