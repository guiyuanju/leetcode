package main

import "fmt"

func main() {
	root := makeBinaryTree([]any{6, 2, 13, 1, 4, 9, 15, nil, nil, nil, nil, nil, nil, 14})
	queries := []int{2, 5, 16}
	fmt.Println(closestNodes(root, queries))

	root = makeBinaryTree([]any{4, nil, 9})
	queries = []int{3}
	fmt.Println(closestNodes(root, queries))
}

func closestNodes(root *TreeNode, queries []int) [][]int {
	var nums []int
	var dfs func(*TreeNode)
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
	bs := func(x int) {
		left := 0
		right := len(nums) - 1
		for left <= right {
			mid := left + (right-left)/2
			if nums[mid] == x {
				res = append(res, []int{x, x})
				return
			} else if nums[mid] > x {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
		l, r := -1, -1
		if right >= 0 {
			l = nums[right]
		}
		if left < len(nums) {
			r = nums[left]
		}
		res = append(res, []int{l, r})
	}

	for _, q := range queries {
		bs(q)
	}

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
