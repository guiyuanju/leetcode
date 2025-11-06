package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}))
	fmt.Println(minPathSum([][]int{{1, 2, 3}, {4, 5, 6}}))
}

func minPathSum(grid [][]int) int {
	// return minPathSumTD(grid)
	return minPathSumBU(grid)
}

func minPathSumTD(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	memo := map[[2]int]int{}
	var dp func(r, c int) int
	dp = func(r, c int) int {
		if v, ok := memo[[2]int{r, c}]; ok {
			return v
		}

		if r == m-1 && c == n-1 {
			return grid[m-1][n-1]
		}

		res := math.MaxInt
		if r+1 < m {
			res = min(res, grid[r][c]+dp(r+1, c))
		}
		if c+1 < n {
			res = min(res, grid[r][c]+dp(r, c+1))
		}

		memo[[2]int{r, c}] = res

		return res
	}

	return dp(0, 0)
}

func minPathSumBU(grid [][]int) int {
	dp := grid
	m := len(grid)
	n := len(grid[0])
	for r := m - 1; r >= 0; r-- {
		for c := n - 1; c >= 0; c-- {
			if r == m-1 && c == n-1 {
				continue
			}
			res := math.MaxInt
			if r+1 < m {
				res = min(res, grid[r][c]+dp[r+1][c])
			}
			if c+1 < n {
				res = min(res, grid[r][c]+dp[r][c+1])
			}
			dp[r][c] = res
		}
	}
	return dp[0][0]
}
