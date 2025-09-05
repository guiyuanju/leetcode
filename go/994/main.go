package main

import (
	"fmt"
)

func main() {
	grid := [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}
	assertEq(4, orangesRotting(grid))

	grid = [][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}
	assertEq(-1, orangesRotting(grid))

	grid = [][]int{{0, 2}}
	assertEq(0, orangesRotting(grid))

	grid = [][]int{{0}}
	assertEq(0, orangesRotting(grid))

	grid = [][]int{{1}}
	assertEq(-1, orangesRotting(grid))

	grid = [][]int{{2}}
	assertEq(0, orangesRotting(grid))
}

func assertEq[T comparable](a, b T) {
	if a != b {
		fmt.Printf("Failed: %v != %v\n", a, b)
	} else {
		fmt.Printf("Ok: %v\n", a)
	}
}

func orangesRotting(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	valid := func(r, c int) bool {
		return 0 <= r && r < m && 0 <= c && c < n && grid[r][c] == 1
	}

	queue := [][2]int{}
	for i := range m {
		for j := range n {
			if grid[i][j] == 2 {
				queue = append(queue, [2]int{i, j})
			}
		}
	}

	step := 0
	for len(queue) > 0 {
		for range len(queue) {
			cur := queue[0]
			queue = queue[1:]
			for _, dir := range directions {
				nextRow, nextCol := cur[0]+dir[0], cur[1]+dir[1]
				if valid(nextRow, nextCol) {
					grid[nextRow][nextCol] = 2
					queue = append(queue, [2]int{nextRow, nextCol})
				}
			}
		}
		if len(queue) > 0 { // !!
			step++
		}
	}

	for i := range m {
		for j := range n {
			if grid[i][j] == 1 {
				return -1
			}
		}
	}

	return step
}
