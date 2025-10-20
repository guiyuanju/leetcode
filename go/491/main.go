package main

import (
	"fmt"
	"slices"
)

func main() {
	nums := []int{4, 6, 7, 7}
	fmt.Println(findSubsequences(nums))

	nums = []int{4, 4, 3, 2, 1}
	fmt.Println(findSubsequences(nums))
}

func findSubsequences(nums []int) [][]int {
	var res [][]int

	var bt func(i int, cur []int)
	bt = func(i int, cur []int) {
		if i > len(nums) {
			return
		}
		if len(cur) >= 2 {
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			res = append(res, tmp)
		}
		for j := i; j < len(nums); j++ {
			if !slices.Contains(nums[i:j], nums[j]) && (len(cur) == 0 || nums[j] >= cur[len(cur)-1]) {
				bt(j+1, append(cur, nums[j]))
			}
		}
	}

	bt(0, nil)

	return res
}
