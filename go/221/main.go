package main

import (
	"github.com/guiyuanju/lcutils/assert"
	"github.com/guiyuanju/lcutils/grid"
)

func main() {
	matrix := grid.Bytes("[[\"1\",\"0\",\"1\",\"0\",\"0\"],[\"1\",\"0\",\"1\",\"1\",\"1\"],[\"1\",\"1\",\"1\",\"1\",\"1\"],[\"1\",\"0\",\"0\",\"1\",\"0\"]]")
	assert.Eq(4, maximalSquare(matrix))

	matrix = grid.Bytes("[[\"0\",\"1\"],[\"1\",\"0\"]]")
	assert.Eq(1, maximalSquare(matrix))

	matrix = grid.Bytes("[[\"0\"]]")
	assert.Eq(0, maximalSquare(matrix))

	matrix = grid.Bytes("[[\"1\",\"1\",\"1\",\"1\",\"0\"],[\"1\",\"1\",\"1\",\"1\",\"0\"],[\"1\",\"1\",\"1\",\"1\",\"1\"],[\"1\",\"1\",\"1\",\"1\",\"1\"],[\"0\",\"0\",\"1\",\"1\",\"1\"]]")
	assert.Eq(16, maximalSquare(matrix))
}

func maximalSquare2(matrix [][]byte) int {
	m := len(matrix)
	n := len(matrix[0])

	var dp func(r, c int) int
	dp = func(r, c int) int {
		if r >= m || c >= n {
			return 0
		}

		if matrix[r][c] == '0' {
			return 0
		}

		length := dp(r+1, c+1)

		if length == 0 {
			return 1
		}

		for i := range length + 1 {
			if matrix[r][c+i] == '0' || matrix[r+i][c] == '0' {
				return i
			}
		}

		return 1 + length
	}

	var res int
	for i := range m {
		for j := range n {
			res = max(res, dp(i, j))
		}
	}

	return res * res
}

func maximalSquare(matrix [][]byte) int {
	m := len(matrix)
	n := len(matrix[0])

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := m - 1; i >= 0; i-- {
	outer:
		for j := n - 1; j >= 0; j-- {
			if matrix[i][j] == '0' {
				dp[i][j] = 0
				continue
			}

			length := dp[i+1][j+1]

			if length == 0 {
				dp[i][j] = 1
				continue
			}

			for k := range length + 1 {
				if matrix[i][j+k] == '0' || matrix[i+k][j] == '0' {
					dp[i][j] = k
					continue outer
				}
			}

			dp[i][j] = 1 + length
		}
	}

	var res int
	for i := range m {
		for j := range n {
			res = max(res, dp[i][j])
		}
	}

	return res * res
}
