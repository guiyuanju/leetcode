package main

import (
	"github.com/guiyuanju/lcutils/assert"
)

func main() {
	assert.Eq(4, longestSubsequence([]int{1, 2, 3, 4}, 1))
	assert.Eq(1, longestSubsequence([]int{1, 3, 5, 7}, 1))
	assert.Eq(4, longestSubsequence([]int{1, 5, 7, 8, 5, 3, 4, 2, 1}, -2))
	assert.Eq(2, longestSubsequence([]int{3, 0, -3, 4, -4, 7, 6}, 3))
	assert.Eq(2, longestSubsequence([]int{4, 12, 10, 0, -2, 7, -8, 9, -9, -12, -12, 8, 8}, 0))
}

func longestSubsequence2(arr []int, difference int) int {
	m := map[int][]int{}
	for i, v := range arr {
		m[v] = append(m[v], i)
	}

	var dp func(i int) int
	dp = func(i int) int {
		if i == 0 {
			return 1
		}

		var res int
		for _, j := range m[arr[i]-difference] {
			if j >= i {
				continue
			}
			res = max(res, dp(j))
		}
		return res + 1
	}

	var res int
	for i := range arr {
		res = max(res, dp(i))
	}

	return res
}

func longestSubsequence3(arr []int, difference int) int {
	dp := make([]int, len(arr))
	dp[0] = 1

	m := map[int][]int{}
	for i, v := range arr {
		m[v] = append(m[v], i)
	}

	var res int
	for i := 1; i < len(arr); i++ {
		for _, j := range m[arr[i]-difference] {
			if j < i {
				dp[i] = dp[j]
			}
		}
		dp[i]++
		res = max(res, dp[i])
	}

	return res
}

func longestSubsequence(arr []int, difference int) int {
	dp := make(map[int]int, len(arr))

	var res int
	for i := range arr {
		if length, ok := dp[arr[i]-difference]; ok {
			dp[arr[i]] = length + 1
		} else {
			dp[arr[i]] = 1
		}
		res = max(res, dp[arr[i]])
	}

	return res
}
