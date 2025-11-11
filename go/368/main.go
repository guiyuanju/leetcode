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
	// return largestDivisibleSubsetTD(nums)
	return largestDivisibleSubsetBU(nums)
}

func largestDivisibleSubsetTD(nums []int) []int {
	memo := map[int][]int{}

	var dp func(i int) []int
	dp = func(i int) []int {
		if v, ok := memo[i]; ok {
			return v
		}

		if i == 0 {
			return []int{nums[0]}
		}

		var res []int
		for j := range i {
			if nums[i]%nums[j] == 0 {
				tmp := dp(j)
				if len(tmp) > len(res) {
					res = tmp
				}
			}
		}
		tmp := make([]int, len(res), len(res)+1)
		copy(tmp, res)
		tmp = append(tmp, nums[i])
		res = tmp

		memo[i] = res

		return res
	}

	slices.Sort(nums)

	var res []int
	for i := range nums {
		tmp := dp(i)
		if len(tmp) > len(res) {
			res = tmp
		}
	}
	return res
}

func largestDivisibleSubsetBU(nums []int) []int {
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

	var length, idx int
	for i := range dp {
		if dp[i] > length {
			length = dp[i]
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
