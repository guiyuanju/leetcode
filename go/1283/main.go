package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{1, 2, 5, 9}
	assertEq(5, smallestDivisor(nums, 6))

	nums = []int{44, 22, 33, 11, 1}
	assertEq(44, smallestDivisor(nums, 5))

	nums = []int{21212, 10101, 12121}
	assertEq(1, smallestDivisor(nums, 1000000))
}

func assertEq[T comparable](a, b T) {
	if a != b {
		fmt.Printf("Failed: %v != %v\n", a, b)
	} else {
		fmt.Printf("Ok: %v\n", a)
	}
}

func smallestDivisor(nums []int, threshold int) int {
	check := func(x int) bool {
		var res int
		for _, n := range nums {
			res += int(math.Ceil(float64(n) / float64(x)))
		}
		return res <= threshold
	}

	var left, right int
	left = 1
	for _, n := range nums {
		right = max(right, n)
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
