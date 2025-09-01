package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(closestValue(makeBinaryTree([]any{4, 2, 5, 1, 3}), 3.714286))
	fmt.Println(closestValue(makeBinaryTree([]any{1}), 4.428571))
}

func closestValue(root *TreeNode, target float64) int {
	var res int
	var dfs func(root *TreeNode, lo, hi int)
	dfs = func(root *TreeNode, lo, hi int) {
		if root == nil {
			if target-float64(lo) <= float64(hi)-target {
				res = lo
			} else {
				res = hi
			}
			return
		}

		if float64(root.Val) == target {
			res = root.Val
		} else if float64(root.Val) > target {
			dfs(root.Left, lo, root.Val)
		} else {
			dfs(root.Right, root.Val, hi)
		}
	}

	dfs(root, math.MinInt, math.MaxInt)

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
