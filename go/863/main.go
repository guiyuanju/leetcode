package main

import "fmt"

func main() {
	fmt.Println(distanceK(makeBinaryTree([]any{3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4}), &TreeNode{5, nil, nil}, 2))
	fmt.Println(distanceK(makeBinaryTree([]any{1}), &TreeNode{1, nil, nil}, 3))
}

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	ups := map[*TreeNode]*TreeNode{}
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		ups[root.Left] = root
		ups[root.Right] = root
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)

	seen := map[*TreeNode]bool{}
	seen[target] = true
	var res []int
	var step int
	q := []*TreeNode{target}
	for len(q) > 0 {
		for range len(q) {
			cur := q[0]
			q = q[1:]
			if step == k {
				res = append(res, cur.Val)
			}
			if cur.Left != nil && !seen[cur.Left] {
				seen[cur.Left] = true
				q = append(q, cur.Left)
			}
			if cur.Right != nil && !seen[cur.Right] {
				seen[cur.Right] = true
				q = append(q, cur.Right)
			}
			if ups[cur] != nil && !seen[ups[cur]] {
				seen[ups[cur]] = true
				q = append(q, ups[cur])
			}
		}
		step++
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
