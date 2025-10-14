package main

import (
	"math"

	"github.com/guiyuanju/lcutils/assert"
)

func main() {
	cookies := []int{8, 15, 10, 20, 8}
	assert.Eq(31, distributeCookies(cookies, 2))

	cookies = []int{6, 1, 3, 2, 2, 4, 1, 2}
	assert.Eq(7, distributeCookies(cookies, 3))
}

func distributeCookies(cookies []int, k int) int {
	res := math.MaxInt
	count := make([]int, k)

	var bt func(i int, hi int)
	bt = func(i int, hi int) {
		if i == len(cookies) {
			res = min(res, hi)
			return
		}

		if hi >= res {
			return
		}

		for j := range k {
			count[j] += cookies[i]
			bt(i+1, max(hi, count[j]))
			count[j] -= cookies[i]
		}
	}

	bt(0, 0)

	return res
}
