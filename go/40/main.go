package main

import (
	"fmt"
	"slices"
)

func main() {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	fmt.Println(combinationSum2(candidates, 8))

	fmt.Println(combinationSum2([]int{2, 5, 2, 1, 2}, 5))

	fmt.Println(combinationSum2([]int{1, 7, 1}, 8))
}

func combinationSum2(candidates []int, target int) [][]int {
	slices.Sort(candidates)

	var res [][]int
	var bt func(i int, sum int, cur []int)
	bt = func(i int, sum int, cur []int) {
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
			if j > i && candidates[j] == candidates[j-1] {
				continue
			}
			bt(j+1, sum+candidates[j], append(cur, candidates[j]))
		}
	}

	bt(0, 0, nil)

	return res
}
