package main

import (
	"fmt"
	"slices"
)

func main() {
	nums := []int{4, 2, 3}
	k := 1
	assertEq(5, largestSumAfterKNegations(nums, k))

	nums = []int{3, -1, 0, 2}
	k = 3
	assertEq(6, largestSumAfterKNegations(nums, k))

	nums = []int{2, -3, -1, 5, -4}
	k = 2
	assertEq(13, largestSumAfterKNegations(nums, k))

	nums = []int{-2, 9, 9, 8, 4}
	k = 5
	assertEq(32, largestSumAfterKNegations(nums, k))

	nums = []int{-8, 3, -5, -3, -5, -2}
	k = 6
	assertEq(22, largestSumAfterKNegations(nums, k))

	nums = []int{-4, -2, -3}
	k = 4
	assertEq(5, largestSumAfterKNegations(nums, k))
}

func assertEq[T comparable](a, b T) {
	if a != b {
		fmt.Printf("Failed: %v != %v\n", a, b)
	} else {
		fmt.Printf("Ok: %v\n", a)
	}
}

func largestSumAfterKNegations(nums []int, k int) int {
	slices.Sort(nums)
	var i int
	for i < len(nums) && k > 0 {
		if nums[i] > 0 {
			break
		}
		nums[i] = -nums[i]
		i++
		k--
	}

	if k%2 == 1 {
		if i == len(nums) {
			nums[i-1] = -nums[i-1]
		} else if i > 0 && nums[i-1] < nums[i] {
			nums[i-1] = -nums[i-1]
		} else {
			nums[i] = -nums[i]
		}
	}

	var res int
	for _, v := range nums {
		res += v
	}
	return res
}
