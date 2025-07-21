package main

import (
	"fmt"
	"slices"
)

func main() {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	fmt.Println(combinationSum2(candidates, 8))

	candidates = []int{2, 5, 2, 1, 2}
	fmt.Println(combinationSum2(candidates, 5))
}

func combinationSum2(candidates []int, target int) [][]int {
	slices.Sort(candidates)
	var res [][]int

	var bt func(cur []int, i int, sum int)
	bt = func(cur []int, i int, sum int) {
		if sum == target {
			res = append(res, append([]int(nil), cur...))
			return
		}

		for j := i; j < len(candidates); j++ {
			if j > i && candidates[j] == candidates[j-1] {
				continue
			}
			newSum := sum + candidates[j]
			if newSum > target {
				break
			}
			bt(append(cur, candidates[j]), j+1, newSum)
		}
	}

	bt(nil, 0, 0)

	return res
}
