package main

import (
	"fmt"
	"reflect"
	"slices"
)

func main() {
	time := []int{1, 2, 3}
	assertEq(int64(3), minimumTime(time, 5))

	time = []int{2}
	assertEq(int64(2), minimumTime(time, 1))
}

func assertEq(a, b any) {
	if reflect.DeepEqual(a, b) {
		fmt.Printf("Ok: %v\n", a)
	} else {
		fmt.Printf("Failed: %v != %v\n", a, b)
	}
}

func minimumTime(time []int, totalTrips int) int64 {
	check := func(total int) bool {
		var res int
		for _, t := range time {
			res += total / t
		}
		return res >= totalTrips
	}

	left := 0
	right := slices.Min(time)*totalTrips + 1
	for left < right {
		mid := left + (right-left)/2
		if check(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return int64(left)
}
