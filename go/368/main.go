package main

import (
	"slices"

	"github.com/guiyuanju/lcutils/assert"
)

func main() {
	assert.Eq([]int{1, 2}, largestDivisibleSubset([]int{1, 2, 3}))
	assert.Eq([]int{1, 2, 4, 8}, largestDivisibleSubset([]int{1, 2, 4, 8}))
	assert.Eq([]int{2, 4, 8}, largestDivisibleSubset([]int{2, 3, 4, 8}))   // 2, 4, 8
	assert.Eq([]int{4, 8, 16}, largestDivisibleSubset([]int{3, 4, 16, 8})) // 4, 8, 16
}

func largestDivisibleSubsetBT(nums []int) []int {
	var res []int
	var bt func(cur []int, i int)
	bt = func(cur []int, i int) {
		if len(cur) > len(res) {
			res = append([]int(nil), cur...)
		}

	outer:
		for j := i; j < len(nums); j++ {
			if len(cur) == 0 {
				bt(append(cur, nums[j]), j+1)
			} else {
				for _, n := range cur {
					if n%nums[j] != 0 && nums[j]%n != 0 {
						continue outer
					}
				}
				bt(append(cur, nums[j]), j+1)
			}
		}
	}

	bt(nil, 0)

	return res
}

func largestDivisibleSubset2(nums []int) []int {
	slices.Sort(nums)
	var dp func(i int) []int
	dp = func(i int) []int {
		if i == 0 {
			return []int{nums[0]}
		}

		var res []int
		for j := range i {
			if nums[i]%nums[j] == 0 {
				tmp := dp(j)
				if len(tmp) > len(res) {
					res = tmp
				}
			}
		}
		return append(res, nums[i])
	}

	var res []int
	for i := range nums {
		tmp := dp(i)
		if len(tmp) > len(res) {
			res = tmp
		}
	}
	return res
}

func largestDivisibleSubset(nums []int) []int {
	slices.Sort(nums)
	dp := make([][]int, len(nums)+1)
	dp[0] = []int{nums[0]}

	for i := 1; i < len(nums); i++ {
		for j := range i {
			if len(dp[j]) > len(dp[i]) && nums[i]%nums[j] == 0 {
				dp[i] = make([]int, len(dp[j]))
				copy(dp[i], dp[j])
			}
		}
		dp[i] = append(dp[i], nums[i])
	}

	var res []int
	for _, v := range dp {
		if len(v) > len(res) {
			res = v
		}
	}
	return res
}
