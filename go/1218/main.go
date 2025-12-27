package main

import (
	"github.com/guiyuanju/lcutils/assert"
)

func main() {
	assert.Eq(4, longestSubsequence([]int{1, 2, 3, 4}, 1))
	assert.Eq(4, longestSubsequence([]int{1, 2, 3, 4}, 1))
	assert.Eq(1, longestSubsequence([]int{1, 3, 5, 7}, 1))
	assert.Eq(4, longestSubsequence([]int{1, 5, 7, 8, 5, 3, 4, 2, 1}, -2))
	assert.Eq(2, longestSubsequence([]int{3, 0, -3, 4, -4, 7, 6}, 3))
	assert.Eq(2, longestSubsequence([]int{4, 12, 10, 0, -2, 7, -8, 9, -9, -12, -12, 8, 8}, 0))
}

func longestSubsequence(arr []int, difference int) int {
	return longestSubsequence_bu(arr, difference)
	// return longestSubsequence_td(arr, difference)
}

func longestSubsequence_td(arr []int, difference int) int {
	m := map[int][]int{}
	for i, n := range arr {
		m[n] = append(m[n], i)
	}

	var dp func(i int) int
	dp = func(i int) int {
		if i == 0 {
			return 1
		}

		var res int
		for _, j := range m[arr[i]-difference] {
			if j < i {
				res = max(res, dp(j))
			}
		}

		return res + 1
	}

	var res int
	for i := range arr {
		res = max(res, dp(i))
	}
	return res
}

// func longestSubsequence(arr []int, difference int) int {
// 	m := make(map[int]int, len(arr))
// 	var res int
// 	for _, n := range arr {
// 		if v, ok := m[n-difference]; ok {
// 			m[n] = v + 1
// 		} else {
// 			m[n] = 1
// 		}
// 		res = max(res, m[n])
// 	}
// 	return res
// }

func longestSubsequence_bu(arr []int, difference int) int {
	m := map[int]int{}
	for _, n := range arr {
		if length, ok := m[n-difference]; ok {
			m[n] = length + 1
		} else {
			m[n] = 1
		}
	}

	var res int
	for _, length := range m {
		res = max(res, length)
	}

	return res
}
