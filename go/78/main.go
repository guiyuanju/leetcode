package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(subsets(nums))

	nums = []int{0}
	fmt.Println(subsets(nums))
}

func subsets(nums []int) [][]int {
	var res [][]int
	var bt func(i int, cur []int)
	bt = func(i int, cur []int) {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		res = append(res, tmp)

		for j := i; j < len(nums); j++ {
			bt(j+1, append(cur, nums[j]))
		}
	}

	bt(0, nil)

	return res
}
