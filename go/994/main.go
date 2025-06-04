package main

import (
	"fmt"
	"math"
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
	M := len(grid)
	N := len(grid[0])
	minutes := make([][]int, M)
	for i := range M {
		minutes[i] = make([]int, N)
		for j := range N {
			minutes[i][j] = math.MaxInt
		}
	}

	type step struct {
		row, col, step int
	}

	valid := func(row, col int) bool {
		return 0 <= row && row < M && 0 <= col && col < N && grid[row][col] == 1
	}

	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	bfs := func(row, col int) {
		seen := map[[2]int]bool{}
		seen[[2]int{row, col}] = true
		queue := []step{{row, col, 0}}
		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]
			for _, dir := range directions {
				newRow, newCol := cur.row+dir[0], cur.col+dir[1]
				if !seen[[2]int{newRow, newCol}] && valid(newRow, newCol) {
					seen[[2]int{newRow, newCol}] = true
					if minutes[newRow][newCol] <= cur.step+1 {
						continue
					}
					minutes[newRow][newCol] = cur.step + 1
					queue = append(queue, step{newRow, newCol, cur.step + 1})
				}
			}
		}
	}

	for i := range M {
		for j := range N {
			if grid[i][j] == 2 {
				bfs(i, j)
			}
		}
	}

	var res int
	for i := range M {
		for j := range N {
			if grid[i][j] == 1 && minutes[i][j] == math.MaxInt {
				return -1
			}
			if minutes[i][j] != math.MaxInt {
				res = max(res, minutes[i][j])
			}
		}
	}
	return res
}
