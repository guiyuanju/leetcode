package main

import (
	"math"

	"github.com/guiyuanju/lcutils/assert"
)

func main() {
	assert.Eq(2, minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
	assert.Eq(1, minSubArrayLen(4, []int{1, 4, 4}))
	assert.Eq(0, minSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1, 1}))
	assert.Eq(5, minSubArrayLen(15, []int{1, 2, 3, 4, 5}))
}

func minSubArrayLen(target int, nums []int) int {
	var i, j, cur int
	res := math.MaxInt
	for ; j < len(nums); j++ {
		cur += nums[j]
		for cur-nums[i] >= target {
			cur -= nums[i]
			i++
		}
		if cur >= target {
			res = min(res, j-i+1)
		}
	}
	if res == math.MaxInt {
		return 0
	}
	return res
}
