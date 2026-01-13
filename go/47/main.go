package main

import (
	"fmt"
	"reflect"
)

func main() {
	nums := []int{1, 1, 2}
	assertEq([][]int{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}}, permuteUnique(nums))

	nums = []int{1, 2, 3}
	assertEq([][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}, permuteUnique(nums))
}

func assertEq(a, b any) {
	if reflect.DeepEqual(a, b) {
		fmt.Printf("Ok: %v\n", a)
	} else {
		fmt.Printf("Failed: %v != %v\n", a, b)
	}
}

func permuteUnique(nums []int) [][]int {
	count := map[int]int{}
	for _, n := range nums {
		count[n]++
	}

	var res [][]int
	var dp func(cur []int)
	dp = func(cur []int) {
		if len(cur) == len(nums) {
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			res = append(res, tmp)
			return
		}

		for n, c := range count {
			if c > 0 {
				count[n]--
				dp(append(cur, n))
				count[n]++
			}
		}
	}

	dp(nil)

	return res
}
