package main

import (
	"fmt"

	. "github.com/guiyuanju/lcutils/binarytree"
)

func main() {
	root := New([]any{3, 2, 3, nil, 3, nil, 1})
	fmt.Println(rob(root))

	root = New([]any{3, 4, 5, 1, 3, nil, 1})
	fmt.Println(rob(root))
}

func rob(root *TreeNode) int {
	memo := map[*TreeNode]int{}
	var dp func(*TreeNode) int
	dp = func(root *TreeNode) int {
		if v, ok := memo[root]; ok {
			return v
		}

		if root == nil {
			return 0
		}

		res := dp(root.Left) + dp(root.Right)
		tmp := root.Val
		if root.Left != nil {
			tmp += dp(root.Left.Left)
			tmp += dp(root.Left.Right)
		}
		if root.Right != nil {
			tmp += dp(root.Right.Left)
			tmp += dp(root.Right.Right)
		}

		res = max(res, tmp)
		memo[root] = res
		return res
	}

	return dp(root)
}
