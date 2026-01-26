package main

import (
	"github.com/guiyuanju/lcutils/assert"
)

func main() {
	assert.Eq(2, numberOfSubarrays([]int{1, 1, 2, 1, 1}, 3))
	assert.Eq(0, numberOfSubarrays([]int{2, 4, 6}, 1))
	assert.Eq(16, numberOfSubarrays([]int{2, 2, 2, 1, 2, 2, 1, 2, 2, 2}, 2))
}

func numberOfSubarrays(nums []int, k int) int {
	count := map[int]int{}
	count[0] = 1
	var res, cur int
	for _, n := range nums {
		if n%2 == 1 {
			cur++
		}
		res += count[cur-k]
		count[cur]++
	}
	return res
}
