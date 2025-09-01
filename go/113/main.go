package main

import "fmt"

func main() {
	root := makeBinaryTree([]any{5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, 5, 1})
	targetSum := 22
	fmt.Println(pathSum(root, targetSum))

	root = makeBinaryTree([]any{1, 2, 3})
	targetSum = 5
	fmt.Println(pathSum(root, targetSum))

	root = makeBinaryTree([]any{1, 2})
	targetSum = 0
	fmt.Println(pathSum(root, targetSum))

	root = makeBinaryTree([]any{-6, nil, -3, -6, 0, -6, -5, 4, nil, nil, nil, 1, 7})
	targetSum = -21
	fmt.Println(pathSum(root, targetSum))
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return nil
	}

	var res [][]int
	var dfs func(root *TreeNode, path []int, curSum int)
	dfs = func(root *TreeNode, path []int, curSum int) {
		if root.Left == nil && root.Right == nil {
			if curSum+root.Val == targetSum {
				tmp := make([]int, len(path)+1)
				copy(tmp, append(path, root.Val))
				res = append(res, tmp)
			}
			return
		}

		if root.Left != nil {
			dfs(root.Left, append(path, root.Val), curSum+root.Val)
		}

		if root.Right != nil {
			dfs(root.Right, append(path, root.Val), curSum+root.Val)
		}
	}

	dfs(root, nil, 0)

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
