package main

import (
	"fmt"
	"reflect"
)

func main() {
	batteries := []int{3, 3, 3}
	assertEq(int64(4), maxRunTime(2, batteries))

	batteries = []int{1, 1, 1, 1}
	assertEq(int64(2), maxRunTime(2, batteries))
}

func assertEq(a, b any) {
	if reflect.DeepEqual(a, b) {
		fmt.Printf("Ok: %v\n", a)
	} else {
		fmt.Printf("Failed: %v != %v\n", a, b)
	}
}

func maxRunTime(n int, batteries []int) int64 {
	check := func(g int64) bool {
		var res int64
		for _, b := range batteries {
			res += min(g, int64(b))
		}
		return res >= g*int64(n)
	}

	var left, right int64
	for _, b := range batteries {
		right += int64(b)
	}
	right /= int64(n)
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
