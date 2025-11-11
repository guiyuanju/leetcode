package main

import (
	"github.com/guiyuanju/lcutils/assert"
)

func main() {
	assert.Eq(2, maxUncrossedLines([]int{1, 4, 2}, []int{1, 2, 4}))
	assert.Eq(3, maxUncrossedLines([]int{2, 5, 1, 2, 5}, []int{10, 5, 2, 1, 5, 2}))
	assert.Eq(2, maxUncrossedLines([]int{1, 3, 7, 1, 7, 5}, []int{1, 9, 2, 5, 1}))
}

func maxUncrossedLines(nums1 []int, nums2 []int) int {
	return maxUncrossedLinesBU(nums1, nums2)
	// return maxUncrossedLinesUD(nums1, nums2)
}

func maxUncrossedLinesUD(nums1 []int, nums2 []int) int {
	memo := map[[2]int]int{}
	var dp func(i, j int) int
	dp = func(i, j int) int {
		if v, ok := memo[[2]int{i, j}]; ok {
			return v
		}

		if i >= len(nums1) || j >= len(nums2) {
			return 0
		}

		var res int
		if nums1[i] == nums2[j] {
			res = 1 + dp(i+1, j+1)
		} else {
			res = max(dp(i+1, j), dp(i, j+1))
		}

		memo[[2]int{i, j}] = res

		return res
	}

	return dp(0, 0)
}

func maxUncrossedLinesBU(nums1 []int, nums2 []int) int {
	dp := make([][]int, len(nums1)+1)
	for i := range dp {
		dp[i] = make([]int, len(nums2)+1)
	}

	for i := len(nums1) - 1; i >= 0; i-- {
		for j := len(nums2) - 1; j >= 0; j-- {
			if nums1[i] == nums2[j] {
				dp[i][j] = 1 + dp[i+1][j+1]
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j+1])
			}
		}
	}

	return dp[0][0]
}
