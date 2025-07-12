package main

import "fmt"

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	assertEq(4, search(nums, 0))

	nums = []int{4, 5, 6, 7, 0, 1, 2}
	assertEq(-1, search(nums, 3))

	nums = []int{1}
	assertEq(-1, search(nums, 0))
}

func assertEq[T comparable](a, b T) {
	if a != b {
		fmt.Printf("Failed: %v != %v\n", a, b)
	} else {
		fmt.Printf("Ok: %v\n", a)
	}
}

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] >= nums[0] {
			left = mid + 1
		} else if nums[mid] <= nums[len(nums)-1] {
			right = mid - 1
		}
	}

	bs := func(nums []int) int {
		left := 0
		right := len(nums) - 1
		for left <= right {
			mid := left + (right-left)/2
			if nums[mid] == target {
				return mid
			} else if nums[mid] > target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
		return -1
	}

	res := bs(nums[:left])
	if res == -1 {
		res = bs(nums[right+1:])
		if res != -1 {
			res += right + 1
		}
	}
	return res
}
