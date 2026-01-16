package main

import (
	"fmt"
)

func main() {
	root := makeBinaryTree([]any{10, 5, -3, 3, 2, nil, 11, 3, -2, nil, 1})
	targetSum := 8
	assertEq(3, pathSum(root, targetSum))

	root = makeBinaryTree([]any{5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, 5, 1})
	targetSum = 22
	assertEq(3, pathSum(root, targetSum))

	root = makeBinaryTree([]any{1})
	targetSum = 0
	assertEq(0, pathSum(root, targetSum))
}

func pathSum(root *TreeNode, targetSum int) int {
	count := map[int]int{}
	count[0] = 1
	var res int
	var dfs func(root *TreeNode, sum int)
	dfs = func(root *TreeNode, sum int) {
		if root == nil {
			return
		}
		sum += root.Val
		res += count[sum-targetSum]
		count[sum]++
		dfs(root.Left, sum)
		dfs(root.Right, sum)
		count[sum]--
	}

	dfs(root, 0)

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
