package main

import "fmt"

func main() {
	fmt.Println(uniquePathsWithObstacles([][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}))
	fmt.Println(uniquePathsWithObstacles([][]int{{0, 1}, {0, 0}}))
	fmt.Println(uniquePathsWithObstacles([][]int{{0, 0}, {0, 1}}))
}

func uniquePathsWithObstacles2(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	if obstacleGrid[m-1][n-1] == 1 {
		return 0
	}

	var dp func(r, c int) int
	dp = func(r, c int) int {
		if r == m-1 && c == n-1 {
			return 1
		}
		if r >= m || c >= n {
			return 0
		}
		if obstacleGrid[r][c] == 1 {
			return 0
		}

		return dp(r+1, c) + dp(r, c+1)
	}

	return dp(0, 0)
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	if obstacleGrid[m-1][n-1] == 1 {
		return 0
	}

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
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
				continue
			}
			dp[i][j] = dp[i+1][j] + dp[i][j+1]
		}
	}

	return dp[0][0]
}
