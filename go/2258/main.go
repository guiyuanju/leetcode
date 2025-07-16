package main

import (
	"github.com/guiyuanju/lcutils/assert"
	"github.com/guiyuanju/lcutils/grid"
)

func main() {
	g := grid.New("[[0,2,0,0,0,0,0],[0,0,0,2,2,1,0],[0,2,0,0,1,2,0],[0,0,2,2,2,0,2],[0,0,0,0,0,0,0]]")
	assert.Eq(3, maximumMinutes(g))

	g = grid.New("[[0,0,0,0],[0,1,2,0],[0,2,0,0]]")
	assert.Eq(-1, maximumMinutes(g))

	g = grid.New("[[0,0,0],[2,2,0],[1,2,0]]")
	assert.Eq(1000000000, maximumMinutes(g))

	g = grid.New("[[0,2,0,0,1],[0,2,0,2,2],[0,2,0,0,0],[0,0,2,2,0],[0,0,0,0,0]]")
	assert.Eq(0, maximumMinutes(g))
}

func maximumMinutes(grid [][]int) int {
	M := len(grid)
	N := len(grid[0])
	directions := [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}
	valid := func(r, c int) bool {
		return 0 <= r && r < M && 0 <= c && c < N && grid[r][c] != 2
	}

	// pre-calc fire spreading
	var q [][2]int
	for i := range M {
		for j := range N {
			if grid[i][j] == 1 {
				grid[i][j] = -1
				q = append(q, [2]int{i, j})
			}
		}
	}
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		for _, dir := range directions {
			nextRow, nextCol := cur[0]+dir[0], cur[1]+dir[1]
			if valid(nextRow, nextCol) && grid[nextRow][nextCol] == 0 {
				grid[nextRow][nextCol] = grid[cur[0]][cur[1]] - 1
				q = append(q, [2]int{nextRow, nextCol})
			}
		}
	}

	// bfs
	seen := make([][]bool, M)
	for i := range seen {
		seen[i] = make([]bool, N)
	}
	check := func(minute int) bool {
		type step struct {
			row, col, minute int
		}
		// clear seen
		for i := range M {
			for j := range N {
				seen[i][j] = false
			}
		}
		q := []step{{0, 0, minute}}
		for len(q) > 0 {
			cur := q[0]
			q = q[1:]
			for _, dir := range directions {
				nextRow, nextCol := cur.row+dir[0], cur.col+dir[1]
				if !valid(nextRow, nextCol) || seen[nextRow][nextCol] {
					continue
				}
				// no fire at save house before moving to it
				if nextRow == M-1 && nextCol == N-1 && (grid[nextRow][nextCol] >= 0 || -grid[nextRow][nextCol]-cur.minute >= 2) {
					return true
				}
				if grid[nextRow][nextCol] >= 0 || // no fire at all time
					-grid[nextRow][nextCol]-cur.minute >= 3 { // no fire after move
					seen[nextRow][nextCol] = true
					q = append(q, step{nextRow, nextCol, cur.minute + 1})
				}
			}
		}
		return false
	}

	// binary search
	left := 0
	right := M*N + 1
	old := right
	for left < right {
		mid := left + (right-left)/2
		if check(mid) {
			left = mid + 1
		} else {
			right = mid
		}
	}
	if left == old {
		return 1e9
	}
	return left - 1
}
