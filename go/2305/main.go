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
	unfairness := make([]int, k)
	var bt func(i int, maximum int, leftChild int)
	bt = func(i int, maximum int, leftChild int) {
		if i == len(cookies) {
			res = maximum
			return
		}

		for j := range k {
			newLeftChild := leftChild
			if unfairness[j] == 0 {
				newLeftChild--
			}
			unfairness[j] += cookies[i]
			newMaximum := max(maximum, unfairness[j])
			if newMaximum >= res || newLeftChild > k-j-1 {
				unfairness[j] -= cookies[i]
				continue
			}
			bt(i+1, newMaximum, newLeftChild)
			unfairness[j] -= cookies[i]
		}
	}

	bt(0, 0, k)

	return res
}
