package main

import (
	"slices"

	"github.com/guiyuanju/lcutils/assert"
)

func main() {
	assert.Eq(4, lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	assert.Eq(4, lengthOfLIS([]int{0, 1, 0, 3, 2, 3}))
	assert.Eq(1, lengthOfLIS([]int{7, 7, 7, 7, 7, 7, 7}))
	assert.Eq(3, lengthOfLIS([]int{4, 10, 4, 3, 8, 9}))
}

func lengthOfLIS2(nums []int) int {
	var dp func(i int) (lenght int)
	dp = func(i int) (lenght int) {
		if i == 0 {
			return 1
		}

		res := 1
		for j := i - 1; j >= 0; j-- {
			if nums[j] < nums[i] {
				res = max(res, dp(j)+1)
			}
		}

		return res
	}

	var res int
	for i := range len(nums) {
		res = max(res, dp(i))
	}
	return res
}

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = 1

	for i := 1; i < len(nums); i++ {
		res := 1
		for j := i - 1; j >= 0; j-- {
			if nums[j] < nums[i] {
				res = max(res, dp[j]+1)
			}
		}
		dp[i] = res
	}

	return slices.Max(dp)
}
