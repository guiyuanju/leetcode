package main

import "fmt"

func main() {
	fmt.Println(uniquePaths(3, 7))
	fmt.Println(uniquePaths(3, 2))
}

func uniquePathsTopDown(m int, n int) int {
	var dp func(r, c int) int
	dp = func(r, c int) int {
		if r == m-1 && c == n-1 {
			return 1
		}
		if r >= m || c >= n {
			return 0
		}

		return dp(r+1, c) + dp(r, c+1)
	}

	return dp(0, 0)
}

func uniquePaths(m int, n int) int {
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	dp[m-1][n-1] = 1

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				continue
			}
			dp[i][j] = dp[i+1][j] + dp[i][j+1]
		}
	}

	return dp[0][0]
}
