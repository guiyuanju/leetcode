package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 1, 2}
	fmt.Println(permuteUnique(nums))

	nums = []int{1, 2, 3}
	fmt.Println(permuteUnique(nums))
}

func permuteUnique(nums []int) [][]int {
	var res [][]int
	count := map[int]int{}
	var bt func(cur []int)
	bt = func(cur []int) {
		if len(cur) == len(nums) {
			res = append(res, append([]int(nil), cur...))
			return
		}

		for n, c := range count {
			if c > 0 {
				count[n]--
				bt(append(cur, n))
				count[n]++
			}
		}
	}

	for _, n := range nums {
		count[n]++
	}
	bt(nil)

	return res
}
