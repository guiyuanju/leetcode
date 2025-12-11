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
	// the end index (exclusive) of prefix sum
	ends := map[int]int{}
	ends[0] = 0
	var cur int
	for i, n := range nums {
		cur += n
		// (cur - x) % k == 0; cur%k -x%k == 0; x%k == cur%k
		if end, ok := ends[cur%k]; ok {
			if i+1-end >= 2 {
				return true
			}
		} else {
			// store the smallest index, satisfy the length greedy
			ends[cur%k] = i + 1
		}
	}
	return false
}
