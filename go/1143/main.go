package main

import (
	"slices"

	"github.com/guiyuanju/lcutils/assert"
)

func main() {
	assert.Eq(3, longestCommonSubsequence("abcde", "ace"))
	assert.Eq(3, longestCommonSubsequence("abc", "abc"))
	assert.Eq(0, longestCommonSubsequence("abc", "def"))
}

func longestCommonSubsequence2(text1 string, text2 string) int {
	var bt func(i, j int) int
	bt = func(i, j int) int {
		if i == 0 {
			if slices.Contains([]byte(text2), text1[0]) {
				return 1
			}
			return 0
		}
		if j == 0 {
			if slices.Contains([]byte(text1), text2[0]) {
				return 1
			}
			return 0
		}

		var res int
		res = bt(i-1, j-1)
		if text1[i] == text2[j] {
			res++
		}
		return max(res, bt(i-1, j), bt(i, j-1))
	}

	return bt(len(text1)-1, len(text2)-1)
}

func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([][]int, len(text1))
	for i := range dp {
		dp[i] = make([]int, len(text2))
	}
	var i int
	for ; i < len(dp[0]); i++ {
		if text1[0] == text2[i] {
			break
		}
	}
	for ; i < len(dp[0]); i++ {
		dp[0][i] = 1
	}
	for i = 0; i < len(dp); i++ {
		if text2[0] == text1[i] {
			break
		}
	}
	for ; i < len(dp); i++ {
		dp[i][0] = 1
	}

	for i = 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			res := dp[i-1][j-1]
			if text1[i] == text2[j] {
				res++
			}
			dp[i][j] = max(res, dp[i][j-1], dp[i-1][j])
		}
	}

	return dp[len(dp)-1][len(dp[0])-1]
}
