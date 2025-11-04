package main

import (
	"slices"

	"github.com/guiyuanju/lcutils/assert"
)

// 1. backtrack
// 2. DP, since there is duplicate computation in searching (if 1 < 2 < 3, no need to seach 3 after 2)
// 3. greedy + bs
// 4. binary indexed tree
// 5. segment tree

func main() {
	assert.Eq(4, lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	assert.Eq(4, lengthOfLIS([]int{0, 1, 0, 3, 2, 3}))
	assert.Eq(1, lengthOfLIS([]int{7, 7, 7, 7, 7, 7, 7}))
	assert.Eq(3, lengthOfLIS([]int{4, 10, 4, 3, 8, 9}))
	assert.Eq(6, lengthOfLIS([]int{1, 3, 6, 7, 9, 4, 10, 5, 6}))
}

func lengthOfLIS(nums []int) int {
	return lengthOfLIS_Greedy(nums)
}

func lengthOfLIS_DP_TD(nums []int) int {
	memo := make([]int, len(nums))
	for i := range memo {
		memo[i] = -1
	}

	var dp func(i int) int
	dp = func(i int) int {
		if memo[i] != -1 {
			return memo[i]
		}

		if i == 0 {
			return 1
		}

		var res int
		for j := range i {
			if nums[j] < nums[i] {
				res = max(res, dp(j))
			}
		}
		res++

		memo[i] = res

		return res
	}

	var res int
	for i := range nums {
		res = max(res, dp(i))
	}
	return res
}

func lengthOfLIS_DP_BU(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = 1
	for i := 1; i < len(nums); i++ {
		var res int
		for j := range i {
			if nums[j] < nums[i] {
				res = max(res, dp[j])
			}
		}
		dp[i] = res + 1
	}
	return slices.Max(dp)
}

func lengthOfLIS_BT(nums []int) int {
	var res int
	var bt func(i int, length int, cur int)
	bt = func(i int, length int, cur int) {
		res = max(res, length)

		for j := i; j < len(nums); j++ {
			if length > 0 && nums[j] > cur || length == 0 {
				bt(j+1, length+1, nums[j])
			}
		}
	}

	bt(0, 0, 0)

	return res
}

// O(n^2), can use binary search -> O(n*long(n))
func lengthOfLIS_Greedy(nums []int) int {
	var cur []int
	for _, n := range nums {
		set := false
		for i, v := range cur {
			if n <= v {
				cur[i] = n
				set = true
				break
			}
		}
		if !set {
			cur = append(cur, n)
		}
	}
	return len(cur)
}
