package main

import (
	"fmt"
	"reflect"
)

func main() {
	bloomDay := []int{1, 10, 3, 10, 2}
	assertEq(3, minDays(bloomDay, 3, 1))

	bloomDay = []int{1, 10, 3, 10, 2}
	assertEq(-1, minDays(bloomDay, 3, 2))

	bloomDay = []int{7, 7, 7, 7, 12, 7, 7}
	assertEq(12, minDays(bloomDay, 2, 3))
}

func assertEq(a, b any) {
	if reflect.DeepEqual(a, b) {
		fmt.Printf("Ok: %v\n", a)
	} else {
		fmt.Printf("Failed: %v != %v\n", a, b)
	}
}

func minDays(bloomDay []int, m int, k int) int {
	check := func(g int) bool {
		var count, adj int
		for _, f := range bloomDay {
			if f > g {
				adj = 0
				continue
			}
			adj++
			if adj < k {
				continue
			}
			count++
			adj = 0
		}
		return count >= m
	}

	left := 1
	right := 0
	for _, f := range bloomDay {
		right = max(right, f)
	}
	right++
	oldRight := right
	for left < right {
		mid := left + (right-left)/2
		if check(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	if left == oldRight {
		return -1
	}
	return left
}
