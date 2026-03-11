package main

import (
	"fmt"
)

func main() {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	fmt.Println(combinationSum2(candidates, 8))

	fmt.Println(combinationSum2([]int{2, 5, 2, 1, 2}, 5))

	fmt.Println(combinationSum2([]int{1, 7, 1}, 8))
}

func combinationSum2(candidates []int, target int) [][]int {
	var res [][]int
	var bt func(i int, cur []int, sum int)
	bt = func(i int, cur []int, sum int) {
		if sum == target {
			tmp := make([]int, i)
			copy(tmp, cur)
			res = append(res, tmp)
		}

		for j := i; j < len(candidates); j++ {
			bt(j+1, append(cur, candidates[j]), sum+candidates[j])
		}
	}

	bt(0, nil, 0)

	return res
}
