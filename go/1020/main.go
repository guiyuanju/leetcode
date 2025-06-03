package main

import "fmt"

func main() {
	grid := [][]int{{0, 0, 0, 0}, {1, 0, 1, 0}, {0, 1, 1, 0}, {0, 0, 0, 0}}
	assertEq(3, numEnclaves(grid))

	grid = [][]int{{0, 1, 1, 0}, {0, 0, 1, 0}, {0, 0, 1, 0}, {0, 0, 0, 0}}
	assertEq(0, numEnclaves(grid))
}

func assertEq[T comparable](a, b T) {
	if a != b {
		fmt.Printf("Failed: %v != %v\n", a, b)
	} else {
		fmt.Printf("Ok: %v\n", a)
	}
}

func numEnclaves(grid [][]int) int {
	const WALK_OFF int = 2

	directions := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	valid := func(row, col int) bool {
		return 0 <= row && row < len(grid) && 0 <= col && col < len(grid[0]) && grid[row][col] == 1
	}
	seen := map[[2]int]bool{}

	var dfs func(r, c int)
	dfs = func(row, col int) {
		for _, dir := range directions {
			nextRow, nextCol := row+dir[0], col+dir[1]
			if !seen[[2]int{nextRow, nextCol}] && valid(nextRow, nextCol) {
				seen[[2]int{nextRow, nextCol}] = true
				grid[nextRow][nextCol] = WALK_OFF
				dfs(nextRow, nextCol)
			}
		}
	}

	var M = len(grid)
	var N = len(grid[0])
	for i := range M {
		dfs(i, -1)
		dfs(i, N)
	}
	for i := range N {
		dfs(-1, i)
		dfs(M, i)
	}

	var res int
	for i := range M {
		for j := range N {
			if grid[i][j] == 1 {
				res++
			}
		}
	}
	return res
}
