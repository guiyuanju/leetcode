package main

import (
	"github.com/guiyuanju/lcutils/assert"
)

func main() {
	assert.Eq(2, maxUncrossedLines([]int{1, 4, 2}, []int{1, 2, 4}))
	assert.Eq(3, maxUncrossedLines([]int{2, 5, 1, 2, 5}, []int{10, 5, 2, 1, 5, 2}))
	assert.Eq(2, maxUncrossedLines([]int{1, 3, 7, 1, 7, 5}, []int{1, 9, 2, 5, 1}))
}

func maxUncrossedLines2(nums1 []int, nums2 []int) int {
	var dp func(i, j int) int
	dp = func(i, j int) int {
		if i < 0 || j < 0 {
			return 0
		}

		if nums1[i] == nums2[j] {
			return 1 + dp(i-1, j-1)
		}
		return max(dp(i-1, j), dp(i, j-1))
	}

	return dp(len(nums1)-1, len(nums2)-1)
}

func maxUncrossedLines(nums1 []int, nums2 []int) int {
	dp := make([][]int, len(nums1)+1)
	for i := range dp {
		dp[i] = make([]int, len(nums2)+1)
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[len(dp)-1][len(dp[0])-1]
}
