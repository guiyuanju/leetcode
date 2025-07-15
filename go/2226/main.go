package main

import (
	"fmt"
	"reflect"
)

func main() {
	candies := []int{5, 8, 6}
	assertEq(5, maximumCandies(candies, 3))

	candies = []int{2, 5}
	assertEq(0, maximumCandies(candies, 11))
}

func assertEq(a, b any) {
	if reflect.DeepEqual(a, b) {
		fmt.Printf("Ok: %v\n", a)
	} else {
		fmt.Printf("Failed: %v != %v\n", a, b)
	}
}

func maximumCandies(candies []int, k int64) int {
	check := func(guess int) bool {
		var count int
		for _, c := range candies {
			count += c / guess
		}
		return int64(count) >= k
	}

	left := 1
	var right int
	for _, c := range candies {
		right = max(right, c)
	}
	right++

	for left < right {
		mid := left + (right-left)/2
		if check(mid) {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left - 1
}
