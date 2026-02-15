package main

import (
	"fmt"
	"reflect"
	"slices"
)

func main() {
	nums := []int{1, 2, 3}
	assertEq([][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}, permute(nums))

	nums = []int{0, 1}
	assertEq([][]int{{0, 1}, {1, 0}}, permute(nums))

	nums = []int{1}
	assertEq([][]int{{1}}, permute(nums))
}

func assertEq(a, b any) {
	if reflect.DeepEqual(a, b) {
		fmt.Printf("Ok: %v\n", a)
	} else {
		fmt.Printf("Failed: %v != %v\n", a, b)
	}
}

func permute(nums []int) [][]int {
	var res [][]int
	var bt func(cur []int)
	bt = func(cur []int) {
		if len(cur) == len(nums) {
			tmp := make([]int, len(nums))
			copy(tmp, cur)
			res = append(res, tmp)
			return
		}

		for i := range nums {
			if !slices.Contains(cur, nums[i]) {
				bt(append(cur, nums[i]))
			}
		}
	}

	bt(nil)

	return res
}
