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
}

func assertEq[T comparable](a, b T) {
	if a != b {
		fmt.Printf("Failed: %v != %v\n", a, b)
	} else {
		fmt.Printf("Ok: %v\n", a)
	}
}

func orangesRotting(grid [][]int) int {
	type step struct {
		row, col, step int
	}

	M := len(grid)
	N := len(grid[0])

	seen := map[[2]int]bool{}
	queue := []step{}
	freshCount := 0
	for i := range M {
		for j := range N {
			if grid[i][j] == 2 {
				queue = append(queue, step{i, j, 0})
				seen[[2]int{i, j}] = true
			} else if grid[i][j] == 1 {
				freshCount++
			}
		}
	}

	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	valid := func(row, col int) bool {
		return 0 <= row && row < M && 0 <= col && col < N && grid[row][col] == 1
	}

	var res int
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		for _, dir := range directions {
			newRow, newCol := cur.row+dir[0], cur.col+dir[1]
			if !seen[[2]int{newRow, newCol}] && valid(newRow, newCol) {
				seen[[2]int{newRow, newCol}] = true
				queue = append(queue, step{newRow, newCol, cur.step + 1})
				res = max(res, cur.step+1)
				freshCount--
			}
		}
	}

	if freshCount > 0 {
		return -1
	}
	return res
}
