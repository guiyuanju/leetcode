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
		return 0 <= r && r < m && 0 <= c && c < n
	}

	var fresh int
	q := [][]int{}
	for i := range m {
		for j := range n {
			if grid[i][j] == 2 {
				q = append(q, []int{i, j})
			} else if grid[i][j] == 1 {
				fresh++
			}
		}
	}

	if fresh == 0 {
		return 0
	}

	var step int
	for len(q) > 0 {
		step++
		for range len(q) {
			cur := q[0]
			q = q[1:]
			for _, dir := range directions {
				nr, nc := cur[0]+dir[0], cur[1]+dir[1]
				if valid(nr, nc) && grid[nr][nc] == 1 {
					grid[nr][nc] = 2
					q = append(q, []int{nr, nc})
					fresh--
					if fresh == 0 {
						return step
					}
				}
			}
		}
	}

	return -1
}
