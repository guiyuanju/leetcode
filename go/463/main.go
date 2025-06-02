package main

import "fmt"

func main() {
	grid := [][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}}
	assertEq(16, islandPerimeter(grid))

	grid = [][]int{{1}}
	assertEq(4, islandPerimeter(grid))

	grid = [][]int{{1, 0}}
	assertEq(4, islandPerimeter(grid))

	grid = [][]int{{0, 1}}
	assertEq(4, islandPerimeter(grid))
}

func assertEq[T comparable](a, b T) {
	if a != b {
		fmt.Printf("Failed: %v != %v\n", a, b)
	} else {
		fmt.Printf("Ok: %v\n", a)
	}
}

func islandPerimeter(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	valid := func(row, col int) bool {
		return 0 <= row && row < m && 0 <= col && col < n
	}
	var res int
	for i := range m {
		for j := range n {
			if grid[i][j] == 0 {
				continue
			}
			for _, dir := range directions {
				nextRow, nextCol := i+dir[0], j+dir[1]
				if !valid(nextRow, nextCol) || grid[nextRow][nextCol] == 0 {
					res++
				}
			}
		}
	}

	return res
}
