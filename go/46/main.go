package main

import (
	"fmt"
	"slices"
)

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(permute(nums))

	nums = []int{0, 1}
	fmt.Println(permute(nums))

	nums = []int{1}
	fmt.Println(permute(nums))
}

func permute(nums []int) [][]int {
	var res [][]int

	var bt func([]int)
	bt = func(cur []int) {
		if len(cur) == len(nums) {
			res = append(res, append([]int(nil), cur...))
			return
		}
		for _, n := range nums {
			if !slices.Contains(cur, n) {
				bt(append(cur, n))
			}
		}
	}

	bt(nil)

	return res
}
