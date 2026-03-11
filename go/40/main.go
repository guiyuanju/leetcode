package main

import (
	"fmt"
	"reflect"
	"slices"
)

func main() {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	assertEq([][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}}, combinationSum2(candidates, 8))

	assertEq([][]int{{1, 2, 2}, {5}}, combinationSum2([]int{2, 5, 2, 1, 2}, 5))

	assertEq([][]int{{1, 7}}, combinationSum2([]int{1, 7, 1}, 8))

	assertEq([][]int{{2}}, combinationSum2([]int{2, 2, 2}, 2))
}

func assertEq(a, b any) {
	if reflect.DeepEqual(a, b) {
		fmt.Printf("Ok: %v\n", a)
	} else {
		fmt.Printf("Failed: %v != %v\n", a, b)
	}
}

func combinationSum2(candidates []int, target int) [][]int {
	slices.Sort(candidates)
	var res [][]int
	var bt func(i int, cur []int, sum int)
	bt = func(i int, cur []int, sum int) {
		if sum == target {
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			res = append(res, tmp)
			return
		}

		for j := i; j < len(candidates); j++ {
			bt(j+1, append(cur, candidates[j]), sum+candidates[j])
			for j < len(candidates)-1 && candidates[j] == candidates[j+1] {
				j++
			}
		}
	}

	bt(0, nil, 0)

	return res
}
