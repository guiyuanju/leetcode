package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}))
	fmt.Println(minPathSum([][]int{{1, 2, 3}, {4, 5, 6}}))
}

func minPathSum2(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	var dp func(r, c int) int
	dp = func(r, c int) int {
		if r == m-1 && c == n-1 {
			return grid[r][c]
		}

		res := math.MaxInt
		if r < m-1 {
			res = min(res, dp(r+1, c))
		}
		if c < n-1 {
			res = min(res, dp(r, c+1))
		}

		return res + grid[r][c]
	}

	return dp(0, 0)
}

func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[m-1][n-1] = grid[m-1][n-1]

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				continue
			}
			res := math.MaxInt
			if i < m-1 {
				res = min(res, dp[i+1][j])
			}
			if j < n-1 {
				res = min(res, dp[i][j+1])
			}
			dp[i][j] = res + grid[i][j]
		}
	}

	return dp[0][0]
}
