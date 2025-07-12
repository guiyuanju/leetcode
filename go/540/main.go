package main

import "fmt"

func main() {
	nums := []int{1, 1, 2, 3, 3, 4, 4, 8, 8}
	assertEq(2, singleNonDuplicate(nums))

	nums = []int{3, 3, 7, 7, 10, 11, 11}
	assertEq(10, singleNonDuplicate(nums))

	nums = []int{1}
	assertEq(1, singleNonDuplicate(nums))
}

func assertEq[T comparable](a, b T) {
	if a != b {
		fmt.Printf("Failed: %v != %v\n", a, b)
	} else {
		fmt.Printf("Ok: %v\n", a)
	}
}

func singleNonDuplicate(nums []int) int {
	check := func(i int) bool {
		if len(nums) <= 1 {
			return true
		}
		if i == 0 {
			return nums[i] != nums[i+1]
		}
		if i == len(nums)-1 {
			return nums[i-1] != nums[i]
		}
		return nums[i-1] != nums[i] && nums[i] != nums[i+1]
	}
	repeatStartIdx := func(i int) int {
		if i == 0 || i == len(nums)-1 {
			return i
		}
		if nums[i-1] == nums[i] {
			return i - 1
		}
		return i
	}
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		if check(mid) {
			return nums[mid]
		} else if repeatStartIdx(mid)%2 == 0 {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
