package main

import (
	"fmt"
	"reflect"
)

func main() {
	root := makeBinaryTree([]any{6, 2, 13, 1, 4, 9, 15, nil, nil, nil, nil, nil, nil, 14})
	queries := []int{2, 5, 16}
	assertEq([][]int{{2, 2}, {4, 6}, {15, -1}}, closestNodes(root, queries))

	root = makeBinaryTree([]any{4, nil, 9})
	queries = []int{3}
	assertEq([][]int{{-1, 4}}, closestNodes(root, queries))
}

func assertEq(a, b any) {
	if reflect.DeepEqual(a, b) {
		fmt.Printf("Ok: %v\n", a)
	} else {
		fmt.Printf("Failed: %v != %v\n", a, b)
	}
}

func closestNodes(root *TreeNode, queries []int) [][]int {
	var nums []int
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		nums = append(nums, root.Val)
		dfs(root.Right)
	}
	dfs(root)

	var res [][]int
	for _, q := range queries {
		idx := bs(nums, q)
		lo, hi := -1, -1
		if idx > 0 {
			lo = nums[idx-1]
		}
		if idx < len(nums) {
			hi = nums[idx]
		}
		if idx < len(nums) && nums[idx] == q {
			lo = q
			hi = q
		}
		res = append(res, []int{lo, hi})
	}

	return res
}

func bs(xs []int, target int) int {
	i := 0
	j := len(xs)
	for i < j {
		mid := i + (j-i)/2
		if xs[mid] < target {
			i = mid + 1
		} else {
			j = mid
		}
	}
	return i
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
