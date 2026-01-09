package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	assertEq(7, maxAncestorDiff(makeBinaryTree([]any{8, 3, 10, 1, 6, nil, 14, nil, nil, 4, 7, 13})))
	assertEq(3, maxAncestorDiff(makeBinaryTree([]any{1, nil, 2, nil, 0, 3})))
	assertEq(5, maxAncestorDiff(makeBinaryTree([]any{2, 4, 3, 1, nil, 0, 5, nil, 6, nil, nil, nil, 7})))
}

func assertEq(a, b any) {
	if reflect.DeepEqual(a, b) {
		fmt.Printf("Ok: %v\n", a)
	} else {
		fmt.Printf("Failed: %v != %v\n", a, b)
	}
}

/*
 *   2
 *  4 3
 * 1  0 5
 *  6     7
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxAncestorDiff(root *TreeNode) int {
	var res int
	var dfs func(root *TreeNode, lo, hi int)
	dfs = func(root *TreeNode, lo, hi int) {
		if root == nil {
			res = max(res, hi-lo)
			return
		}
		dfs(root.Left, min(lo, root.Val), max(hi, root.Val))
		dfs(root.Right, min(lo, root.Val), max(hi, root.Val))
	}

	dfs(root, math.MaxInt, math.MinInt)

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
