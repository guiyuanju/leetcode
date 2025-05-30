package main

import "fmt"

func main() {
	root1 := makeBinaryTree([]any{2, 1, 4})
	root2 := makeBinaryTree([]any{1, 0, 3})
	fmt.Println(getAllElements(root1, root2))

	root1 = makeBinaryTree([]any{1, nil, 8})
	root2 = makeBinaryTree([]any{8, 1})
	fmt.Println(getAllElements(root1, root2))

}

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	var res []int
	a1 := flatten(root1)
	a2 := flatten(root2)
	var i, j int
	for i < len(a1) && j < len(a2) {
		if a1[i] < a2[j] {
			res = append(res, a1[i])
			i++
		} else {
			res = append(res, a2[j])
			j++
		}
	}
	if i < len(a1) {
		res = append(res, a1[i:]...)
	}
	if j < len(a2) {
		res = append(res, a2[j:]...)
	}
	return res
}

func flatten(root *TreeNode) []int {
	var res []int
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		res = append(res, root.Val)
		dfs(root.Right)
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
