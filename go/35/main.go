package main

import "fmt"

func main() {
	nums := []int{1, 3, 5, 6}
	target := 5
	assertEq(2, searchInsert(nums, target))

	nums = []int{1, 3, 5, 6}
	target = 2
	assertEq(1, searchInsert(nums, target))

	nums = []int{1, 3, 5, 6}
	target = 7
	assertEq(4, searchInsert(nums, target))
}

func assertEq[T comparable](a, b T) {
	if a != b {
		fmt.Printf("Failed: %v != %v\n", a, b)
	} else {
		fmt.Printf("Ok: %v\n", a)
	}
}

func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums)
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
