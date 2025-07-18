package main

import (
	"github.com/guiyuanju/lcutils/assert"
	"github.com/guiyuanju/lcutils/grid"
)

func main() {
	candidates := []int{2, 3, 6, 7}
	assert.Eq(grid.New("[[2,2,3],[7]]"), combinationSum(candidates, 7))

	candidates = []int{2, 3, 5}
	assert.Eq(grid.New("[[2,2,2,2],[2,3,3],[3,5]]"), combinationSum(candidates, 8))
}

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var bt func(cur []int, sum int, i int)
	bt = func(cur []int, sum int, i int) {
		if sum == target {
			res = append(res, append([]int(nil), cur...))
			return
		}
		if sum > target {
			return
		}
		for j := i; j < len(candidates); j++ {
			bt(append(cur, candidates[j]), sum+candidates[j], j)
		}
	}
	bt(nil, 0, 0)
	return res
}
