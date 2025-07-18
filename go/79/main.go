package main

import (
	"github.com/guiyuanju/lcutils/assert"
	"github.com/guiyuanju/lcutils/grid"
)

func main() {
	board := grid.Bytes("[[\"A\",\"B\",\"C\",\"E\"],[\"S\",\"F\",\"C\",\"S\"],[\"A\",\"D\",\"E\",\"E\"]]")
	word := "SEE"
	assert.Eq(true, exist([][]byte(board), word))

	board = grid.Bytes("[[\"A\",\"B\",\"C\",\"E\"],[\"S\",\"F\",\"C\",\"S\"],[\"A\",\"D\",\"E\",\"E\"]]")
	word = "ABCB"
	assert.Eq(false, exist([][]byte(board), word))
}

func exist(board [][]byte, word string) bool {
	visited := make([][]bool, len(board))
	for i := range visited {
		visited[i] = make([]bool, len(board[0]))
	}

	valid := func(r, c int) bool {
		return 0 <= r && r < len(board) && 0 <= c && c < len(board[0])
	}
	directions := [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
	var bt func(r, c, i int) bool
	bt = func(r, c, i int) bool {
		if i == len(word) {
			return true
		}
		for _, dir := range directions {
			nextRow, nextCol := r+dir[0], c+dir[1]
			if valid(nextRow, nextCol) && !visited[nextRow][nextCol] && board[nextRow][nextCol] == word[i] {
				visited[nextRow][nextCol] = true
				if bt(nextRow, nextCol, i+1) {
					return true
				}
				visited[nextRow][nextCol] = false
			}
		}
		return false
	}

	for i := range len(board) {
		for j := range len(board[0]) {
			if board[i][j] == word[0] {
				visited[i][j] = true
				if bt(i, j, 1) {
					return true
				}
				visited[i][j] = false
			}
		}
	}

	return false
}
