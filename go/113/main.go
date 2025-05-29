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
	var res [][]int

	var dfs func(*TreeNode, int, []int)
	dfs = func(root *TreeNode, remain int, path []int) {
		if root == nil {
			return
		}
		if root.Left == nil && root.Right == nil {
			if remain == root.Val {
				path = append(path, root.Val)
				tmp := make([]int, len(path))
				copy(tmp, path)
				res = append(res, tmp)
			}
			return
		}

		path = append(path, root.Val)
		dfs(root.Left, remain-root.Val, path)
		dfs(root.Right, remain-root.Val, path)
	}

	dfs(root, targetSum, []int{})
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
