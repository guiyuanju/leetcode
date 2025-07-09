package main

import "fmt"

func main() {
	nums := []int{7, 2, 5, 10, 8}
	assertEq(18, splitArray(nums, 2))

	nums = []int{1, 2, 3, 4, 5}
	assertEq(9, splitArray(nums, 2))
}

func assertEq[T comparable](a, b T) {
	if a != b {
		fmt.Printf("Failed: %v != %v\n", a, b)
	} else {
		fmt.Printf("Ok: %v\n", a)
	}
}

func splitArray(nums []int, k int) int {
	check := func(value int) bool {
		var cur, count int
		for _, n := range nums {
			if n > value {
				return false
			}
			cur += n
			if cur > value {
				count++
				cur = n
			}
		}
		return count+1 <= k
	}

	var left, right int
	for _, n := range nums {
		right += n
	}

	for left <= right {
		mid := left + (right-left)/2
		if check(mid) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}
