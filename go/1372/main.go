package main

import "fmt"

func main() {
	root := makeBinaryTree([]any{1, nil, 1, 1, 1, nil, nil, 1, 1, nil, 1, nil, nil, nil, 1})
	assertEq(3, longestZigZag(root))

	root = makeBinaryTree([]any{1, 1, 1, nil, 1, nil, nil, 1, 1, nil, 1})
	assertEq(4, longestZigZag(root))

	root = makeBinaryTree([]any{1})
	assertEq(0, longestZigZag(root))
}

func longestZigZag(root *TreeNode) int {
	var res int
	var dfs func(root *TreeNode, cur int, left bool)
	dfs = func(root *TreeNode, cur int, left bool) {
		if root == nil {
			res = max(res, cur)
			return
		}

		if left {
			dfs(root.Left, 0, true)
			dfs(root.Right, cur+1, false)
		} else {
			dfs(root.Left, cur+1, true)
			dfs(root.Right, 0, false)
		}
	}

	dfs(root.Left, 0, true)
	dfs(root.Right, 0, false)

	return res
}

func assertEq[T comparable](a, b T) {
	if a != b {
		fmt.Printf("Failed: %v != %v\n", a, b)
	} else {
		fmt.Printf("Ok: %v\n", a)
	}
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
