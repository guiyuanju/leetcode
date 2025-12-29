package main

import (
	"slices"

	"github.com/guiyuanju/lcutils/assert"
)

func main() {
	assert.Eq([]int{1, 2}, largestDivisibleSubset([]int{1, 2, 3}))
	assert.Eq([]int{1, 2, 4, 8}, largestDivisibleSubset([]int{1, 2, 4, 8}))
	assert.Eq([]int{2, 4, 8}, largestDivisibleSubset([]int{2, 3, 4, 8}))   // 2, 4, 8
	assert.Eq([]int{4, 8, 16}, largestDivisibleSubset([]int{3, 4, 16, 8})) // 4, 8, 16
}

func largestDivisibleSubset(nums []int) []int {
	// return largestDivisibleSubset_td(nums)
	// return largestDivisibleSubset_bu(nums)
	return largestDivisibleSubset_bu_space_opt(nums)
}

func largestDivisibleSubset_td(nums []int) []int {
	slices.Sort(nums)
	var dp func(i int) []int
	dp = func(i int) []int {
		if i == 0 {
			return []int{nums[0]}
		}

		var res []int
		for j := range i {
			if nums[i]%nums[j] == 0 {
				if cur := dp(j); len(cur) > len(res) {
					res = cur
				}
			}
		}

		return append(res, nums[i])
	}

	var res []int
	for i := range nums {
		if cur := dp(i); len(cur) > len(res) {
			res = cur
		}
	}

	return res
}

func largestDivisibleSubset_bu(nums []int) []int {
	slices.Sort(nums)
	dp := make([][]int, len(nums))
	dp[0] = []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		for j := range i {
			if nums[i]%nums[j] == 0 && len(dp[j]) > len(dp[i]) {
				dp[i] = dp[j]
			}
		}
		dp[i] = append(dp[i], nums[i])
	}

	var res []int
	for _, v := range dp {
		if len(v) > len(res) {
			res = v
		}
	}

	return res
}

func largestDivisibleSubset_bu_space_opt(nums []int) []int {
	slices.Sort(nums)
	dp := make([]int, len(nums))
	dp[0] = 1
	prev := make([]int, len(nums))
	for i := range prev {
		prev[i] = -1
	}

	for i := 1; i < len(nums); i++ {
		for j := range i {
			if nums[i]%nums[j] == 0 && dp[j] > dp[i] {
				dp[i] = dp[j]
				prev[i] = j
			}
		}
		dp[i]++
	}

	var idx, hi int
	for i, v := range dp {
		if v > hi {
			hi = v
			idx = i
		}
	}

	var res []int
	for idx != -1 {
		res = append(res, nums[idx])
		idx = prev[idx]
	}

	slices.Reverse(res)

	return res
}
