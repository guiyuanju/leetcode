package main

import (
	"fmt"
	"math"
)

func main() {
	matrix := [][]int{{2, 1, 3}, {6, 5, 4}, {7, 8, 9}}
	fmt.Println(minFallingPathSum(matrix))

	matrix = [][]int{{-19, 57}, {-40, -5}}
	fmt.Println(minFallingPathSum(matrix))
}

func minFallingPathSum2(matrix [][]int) int {
	m := len(matrix)
	n := len(matrix[0])

	var dp func(r, c int) int
	dp = func(r, c int) int {
		if r == m-1 {
			return matrix[r][c]
		}

		res := math.MaxInt
		if c > 0 {
			res = min(res, dp(r+1, c-1))
		}
		res = min(res, dp(r+1, c))
		if c < n-1 {
			res = min(res, dp(r+1, c+1))
		}
		return res + matrix[r][c]
	}

	res := math.MaxInt
	for i := range n {
		res = min(res, dp(0, i))
	}
	return res
}

func minFallingPathSum3(matrix [][]int) int {
	m := len(matrix)
	n := len(matrix[0])

	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 {
				dp[i][j] = matrix[i][j]
				continue
			}
			res := math.MaxInt
			if j > 0 {
				res = min(res, dp[i+1][j-1])
			}
			res = min(res, dp[i+1][j])
			if j < n-1 {
				res = min(res, dp[i+1][j+1])
			}
			dp[i][j] = res + matrix[i][j]
		}
	}

	res := math.MaxInt
	for i := range n {
		res = min(res, dp[0][i])
	}
	return res
}

func minFallingPathSum(matrix [][]int) int {
	m := len(matrix)
	n := len(matrix[0])

	dp := make([]int, m)
	tmp := make([]int, m)

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 {
				dp[j] = matrix[i][j]
				continue
			}
			res := math.MaxInt
			if j > 0 {
				res = min(res, dp[j-1])
			}
			res = min(res, dp[j])
			if j < n-1 {
				res = min(res, dp[j+1])
			}
			tmp[j] = res + matrix[i][j]
			if j == 0 {
				dp, tmp = tmp, dp
			}
		}
	}

	res := math.MaxInt
	for i := range n {
		res = min(res, dp[i])
	}
	return res
}
