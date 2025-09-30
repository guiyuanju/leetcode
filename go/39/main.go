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
	var bt func(i int, cur []int, sum int)
	bt = func(i int, cur []int, sum int) {
		if sum > target {
			return
		}
		if sum == target {
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			res = append(res, tmp)
			return
		}

		for j := i; j < len(candidates); j++ {
			bt(j, append(cur, candidates[j]), sum+candidates[j])
		}
	}

	bt(0, nil, 0)

	return res
}
