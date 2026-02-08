package main

import (
	"fmt"
	"reflect"
)

func main() {
	root := makeBinaryTree([]any{6, 2, 8, 0, 4, 7, 9, nil, nil, 3, 5})
	p := &TreeNode{2, nil, nil}
	q := &TreeNode{8, nil, nil}
	assertEq(6, lowestCommonAncestor(root, p, q).Val)

	root = makeBinaryTree([]any{6, 2, 8, 0, 4, 7, 9, nil, nil, 3, 5})
	p = &TreeNode{2, nil, nil}
	q = &TreeNode{4, nil, nil}
	assertEq(2, lowestCommonAncestor(root, p, q).Val)

	root = makeBinaryTree([]any{2, 1})
	p = &TreeNode{2, nil, nil}
	q = &TreeNode{1, nil, nil}
	assertEq(2, lowestCommonAncestor(root, p, q).Val)
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	for root != nil {
		if p.Val < root.Val && q.Val < root.Val {
			root = root.Left
		} else if p.Val > root.Val && q.Val > root.Val {
			root = root.Right
		} else {
			return root
		}
	}

	return nil
}

func pp(nodes []*TreeNode) {
	for _, n := range nodes {
		fmt.Print(n.Val, " ")
	}
	fmt.Println()
}

func assertEq(a, b any) {
	if reflect.DeepEqual(a, b) {
		fmt.Printf("Ok: %v\n", a)
	} else {
		fmt.Printf("Failed: %v != %v\n", a, b)
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
