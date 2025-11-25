package main

import (
	"github.com/guiyuanju/lcutils/assert"
)

func main() {
	assert.Eq(false, carPooling([][]int{{2, 1, 5}, {3, 3, 7}}, 4))
	assert.Eq(true, carPooling([][]int{{2, 1, 5}, {3, 3, 7}}, 5))
	assert.Eq(true, carPooling([][]int{{3, 2, 9}, {3, 2, 8}, {5, 2, 5}, {7, 4, 8}, {8, 7, 8}, {9, 1, 2}}, 22))
}

func carPooling(trips [][]int, capacity int) bool {
	diff := make([]int, 1001)
	hi := 0
	for _, t := range trips {
		diff[t[1]] += t[0]
		diff[t[2]] -= t[0]
		hi = max(hi, t[1], t[2])
	}

	var cur int
	for i := range hi + 1 {
		cur += diff[i]
		if cur > capacity {
			return false
		}
	}

	return true
}
