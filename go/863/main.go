package main

import "fmt"

func main() {
	tree := makeBinaryTree([]any{3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4})
	fmt.Println(distanceK(tree, getNodeFromUniqVal(tree, 5), 2))
	fmt.Println(distanceK(makeBinaryTree([]any{1}), &TreeNode{1, nil, nil}, 3))
}

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	up := map[int]*TreeNode{}
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Left != nil {
			up[root.Left.Val] = root
			dfs(root.Left)
		}
		if root.Right != nil {
			up[root.Right.Val] = root
			dfs(root.Right)
		}
	}

	dfs(root)

	q := []*TreeNode{target}
	seen := map[int]bool{}
	seen[target.Val] = true
	var step int
	var res []int
	for len(q) > 0 {
		for range len(q) {
			cur := q[0]
			q = q[1:]
			if step == k {
				res = append(res, cur.Val)
			}
			if cur.Left != nil && !seen[cur.Left.Val] {
				q = append(q, cur.Left)
				seen[cur.Left.Val] = true
			}
			if cur.Right != nil && !seen[cur.Right.Val] {
				q = append(q, cur.Right)
				seen[cur.Right.Val] = true
			}
			if up[cur.Val] != nil && !seen[up[cur.Val].Val] {
				q = append(q, up[cur.Val])
				seen[up[cur.Val].Val] = true
			}
		}
		if step == k {
			break
		}
		step++
	}

	return res
}

// Suppose the tree node has unique value, return the node of provided value
func getNodeFromUniqVal(root *TreeNode, val int) *TreeNode {
	var dp func(root *TreeNode) *TreeNode
	dp = func(root *TreeNode) *TreeNode {
		if root == nil {
			return nil
		}
		if root.Val == val {
			return root
		}
		res := dp(root.Left)
		if res == nil {
			res = dp(root.Right)
		}
		return res
	}

	return dp(root)
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
