package main

import "fmt"

func main() {
	root := makeBinaryTree([]any{5, 3, 6, 2, 4, nil, 7})
	key := 3
	printBinaryTree(deleteNode(root, key))

	root = makeBinaryTree([]any{5, 3, 6, 2, 4, nil, 7})
	key = 0
	printBinaryTree(deleteNode(root, key))

	root = makeBinaryTree([]any{})
	key = 0
	printBinaryTree(deleteNode(root, key))
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == key {
		return insertRightMost(root.Left, root.Right)
	}

	root.Left = deleteNode(root.Left, key)
	root.Right = deleteNode(root.Right, key)

	return root
}

func insertRightMost(root *TreeNode, n *TreeNode) *TreeNode {
	if root == nil {
		return n
	}
	if n == nil {
		return root
	}
	cur := root
	for cur != nil {
		if cur.Right == nil {
			cur.Right = n
			return root
		}
		cur = cur.Right
	}
	return root
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
