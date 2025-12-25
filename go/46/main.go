package main

import (
	"fmt"
	"reflect"
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
	m := make(map[int]bool, len(nums))
	cur := make([]int, len(nums))
	var bt func(i int)
	bt = func(i int) {
		if i == len(nums) {
			tmp := make([]int, len(nums))
			copy(tmp, cur)
			res = append(res, tmp)
			return
		}
		for _, n := range nums {
			if !m[n] {
				m[n] = true
				cur[i] = n
				bt(i + 1)
				m[n] = false
			}
		}
	}

	bt(0)

	return res
}
