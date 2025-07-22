package main

import (
	"github.com/guiyuanju/lcutils/assert"
	"github.com/guiyuanju/lcutils/grid"
)

func main() {
	g := grid.Ints("[[1,0,0,0],[0,0,0,0],[0,0,2,-1]]")
	assert.Eq(2, uniquePathsIII(g))

	g = grid.Ints("[[1,0,0,0],[0,0,0,0],[0,0,0,2]]")
	assert.Eq(4, uniquePathsIII(g))

	g = grid.Ints("[[0,1],[2,0]]")
	assert.Eq(0, uniquePathsIII(g))
}

func uniquePathsIII(grid [][]int) int {
	M := len(grid[0])
	N := len(grid)
	seen := make([][]bool, N)
	for i := range seen {
		seen[i] = make([]bool, M)
	}

	var emptyBlocks int
	for i := range N {
		for j := range M {
			if grid[i][j] == 0 {
				emptyBlocks++
			}
		}
	}

	directions := [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}
	valid := func(r, c int) bool {
		return 0 <= r && r < N && 0 <= c && c < M && grid[r][c] != -1
	}

	var res int
	var dfs func(r, c, steps int)
	dfs = func(r, c, steps int) {
		if grid[r][c] == 2 {
			if steps-1 == emptyBlocks {
				res++
			}
			return
		}

		for _, dir := range directions {
			nextRow, nextCol := r+dir[0], c+dir[1]
			if valid(nextRow, nextCol) && !seen[nextRow][nextCol] {
				seen[nextRow][nextCol] = true
				dfs(nextRow, nextCol, steps+1)
				seen[nextRow][nextCol] = false
			}
		}
	}

	for i := range N {
		for j := range M {
			if grid[i][j] == 1 {
				seen[i][j] = true
				dfs(i, j, 0)
				return res
			}
		}
	}

	return res
}
