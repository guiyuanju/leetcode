package main

import "fmt"

func main() {
	fmt.Println(distanceK(makeBinaryTree([]any{3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4}), &TreeNode{5, nil, nil}, 2))
	fmt.Println(distanceK(makeBinaryTree([]any{1}), &TreeNode{1, nil, nil}, 3))
}

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	g := map[int][]int{}
	var toGraph func(root *TreeNode, paren *TreeNode)
	toGraph = func(root *TreeNode, paren *TreeNode) {
		if root == nil {
			return
		}
		if paren != nil {
			g[root.Val] = append(g[root.Val], paren.Val)
			g[paren.Val] = append(g[paren.Val], root.Val)
		}
		toGraph(root.Left, root)
		toGraph(root.Right, root)
	}
	toGraph(root, nil)

	seen := map[int]bool{}
	queue := []int{target.Val}
	seen[target.Val] = true
	for k > 0 && len(queue) > 0 {
		k--
		for range len(queue) {
			cur := queue[0]
			queue = queue[1:]
			for _, nei := range g[cur] {
				if !seen[nei] {
					seen[nei] = true
					queue = append(queue, nei)
				}
			}
		}
	}

	return queue
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
