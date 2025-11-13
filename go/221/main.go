package main

import (
	"github.com/guiyuanju/lcutils/assert"
	"github.com/guiyuanju/lcutils/grid"
)

func main() {
	matrix := grid.Bytes("[[\"1\",\"0\",\"1\",\"0\",\"0\"],[\"1\",\"0\",\"1\",\"1\",\"1\"],[\"1\",\"1\",\"1\",\"1\",\"1\"],[\"1\",\"0\",\"0\",\"1\",\"0\"]]")
	assert.Eq(4, maximalSquare(matrix))

	assert.Eq(1, maximalSquare(grid.Bytes("[[\"0\",\"1\"],[\"1\",\"0\"]]")))

	matrix = grid.Bytes("[[\"0\"]]")
	assert.Eq(0, maximalSquare(matrix))

	matrix = grid.Bytes("[[\"1\",\"1\",\"1\",\"1\",\"0\"],[\"1\",\"1\",\"1\",\"1\",\"0\"],[\"1\",\"1\",\"1\",\"1\",\"1\"],[\"1\",\"1\",\"1\",\"1\",\"1\"],[\"0\",\"0\",\"1\",\"1\",\"1\"]]")
	assert.Eq(16, maximalSquare(matrix))
}

func maximalSquare(matrix [][]byte) int {
	return maximalSquareTD(matrix)
}

func maximalSquareTD(matrix [][]byte) int {
	m := len(matrix)
	n := len(matrix[0])

	memo := map[[2]int]int{}

	var dp func(r, c int) int
	dp = func(r, c int) int {
		if v, ok := memo[[2]int{r, c}]; ok {
			return v
		}

		if r >= m || c >= n || matrix[r][c] == '0' {
			return 0
		}

		res := 1 + min(dp(r+1, c+1), dp(r+1, c), dp(r, c+1))
		memo[[2]int{r, c}] = res
		return res
	}

	var res int
	for i := range m {
		for j := range n {
			res = max(res, dp(i, j))
		}
	}
	return res * res
}
