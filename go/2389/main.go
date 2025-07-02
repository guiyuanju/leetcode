package main

import (
	"fmt"
	"slices"
)

func main() {
	nums := []int{4, 5, 2, 1}
	queries := []int{3, 10, 21}
	fmt.Println(answerQueries(nums, queries))

	nums = []int{2, 3, 4, 5}
	queries = []int{1}
	fmt.Println(answerQueries(nums, queries))
}

func answerQueries(nums []int, queries []int) []int {
	slices.Sort(nums)
	presum := make([]int, len(nums))
	presum[0] = nums[0]
	for i := 1; i < len(presum); i++ {
		presum[i] = presum[i-1] + nums[i]
	}

	var res []int
	for _, q := range queries {
		idx := bs(presum, q)
		if idx < len(presum) && presum[idx] == q {
			res = append(res, idx+1)
		} else {
			res = append(res, idx)
		}
	}
	return res
}

func bs(nums []int, target int) int {
	left := 0
	right := len(nums)
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
