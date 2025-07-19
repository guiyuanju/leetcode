package main

import (
	"github.com/guiyuanju/lcutils/assert"
	"github.com/guiyuanju/lcutils/grid"
	"github.com/guiyuanju/lcutils/math"
)

func main() {
	assert.Eq(grid.Ints("[[1,2,4]]"), combinationSum3(3, 7))
	assert.Eq(grid.Ints("[[1,2,6],[1,3,5],[2,3,4]]"), combinationSum3(3, 9))
	assert.Eq(grid.Ints("[]"), combinationSum3(4, 1))
}

func combinationSum3(k int, n int) [][]int {
	var res [][]int

	var bt func(cur []int, i int)
	bt = func(cur []int, i int) {
		if len(cur) == k {
			if math.Sum(cur) == n {
				res = append(res, append([]int(nil), cur...))
			}
			return
		}

		for j := i; j < 10; j++ {
			bt(append(cur, j), j+1)
		}
	}

	bt(nil, 1)
	return res
}
