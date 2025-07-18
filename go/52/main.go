package main

import "github.com/guiyuanju/lcutils/assert"

func main() {
	assert.Eq(2, totalNQueens(4))
	assert.Eq(1, totalNQueens(1))
}

func totalNQueens(n int) int {
	valid := func(exist []int, col int) bool {
		row := len(exist)
		for r, c := range exist {
			if c == col || abs(col-c) == row-r {
				return false
			}
		}
		return true
	}

	var res int
	var bt func(exist []int)
	bt = func(exist []int) {
		if len(exist) >= n {
			res++
			return
		}
		for i := range n {
			if valid(exist, i) {
				bt(append(exist, i))
			}
		}
	}

	bt(nil)
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
