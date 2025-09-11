package main

import (
	"fmt"
	"slices"
)

func main() {
	nums := []int{3, 6, 1, 2, 5}
	k := 2
	assertEq(2, (partitionArray(nums, k)))

	nums = []int{1, 2, 3}
	k = 1
	assertEq(2, (partitionArray(nums, k)))

	nums = []int{2, 2, 4, 5}
	k = 0
	assertEq(3, (partitionArray(nums, k)))

	nums = []int{3, 1, 3, 4, 2}
	k = 0
	assertEq(4, (partitionArray(nums, k)))
}

func assertEq[T comparable](a, b T) {
	if a != b {
		fmt.Printf("Failed: %v != %v\n", a, b)
	} else {
		fmt.Printf("Ok: %v\n", a)
	}
}

func partitionArray(nums []int, k int) int {
	slices.Sort(nums)
	res := 1
	start := nums[0]
	for _, n := range nums {
		if n-start > k {
			start = n
			res++
		}
	}
	return res
}
