package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(getMinimumDifference(makeBinaryTree([]any{4, 2, 6, 1, 3})))
	fmt.Println(getMinimumDifference(makeBinaryTree([]any{1, 0, 48, nil, nil, 12, 49})))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func getMinimumDifference(root *TreeNode) int {
	prev := -1
	res := math.MaxInt
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		dfs(root.Left)
		if prev != -1 {
			res = min(res, abs(prev-root.Val))
		}
		prev = root.Val
		dfs(root.Right)
	}
	dfs(root)
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
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
