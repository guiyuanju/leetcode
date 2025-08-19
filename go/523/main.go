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
}

// (cur-sum)%k == 0
// cur%k - sum%k == 0
// sum%k == cur%k
func checkSubarraySum(nums []int, k int) bool {
	if len(nums) < 2 {
		return false
	}

	sums := map[int]int{}
	sums[0] = -1
	var cur int
	for j := 0; j < len(nums); j++ {
		cur += nums[j]
		if v, ok := sums[cur%k]; ok {
			if j-v > 1 {
				return true
			}
		} else {
			sums[cur%k] = j
		}
	}
	return false
}
