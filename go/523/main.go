package main

import (
	"github.com/guiyuanju/lcutils/assert"
)

func main() {
	assert.Eq(true, checkSubarraySum([]int{23, 2, 4, 6, 7}, 6))
	assert.Eq(true, checkSubarraySum([]int{23, 2, 6, 4, 7}, 6))
	assert.Eq(false, checkSubarraySum([]int{23, 2, 6, 4, 7}, 13))
	assert.Eq(false, checkSubarraySum([]int{1, 0}, 2))
	assert.Eq(true, checkSubarraySum([]int{5, 0, 0, 0}, 3))
	assert.Eq(false, checkSubarraySum([]int{1, 2, 12}, 6))
	assert.Eq(false, checkSubarraySum([]int{1}, 1))
}

func checkSubarraySum(nums []int, k int) bool {
	var cur int
	sum := map[int]int{}
	sum[0] = -1
	for i, n := range nums {
		cur += n
		if idx, ok := sum[cur%k]; ok {
			if i-idx >= 2 {
				return true
			}
		} else {
			sum[cur%k] = i
		}
	}
	return false
}
