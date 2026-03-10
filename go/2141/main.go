package main

import (
	"fmt"
	"reflect"
	"slices"
)

func main() {
	batteries := []int{3, 3, 3}
	assertEq(int64(4), maxRunTime(2, batteries))

	batteries = []int{1, 1, 1, 1}
	assertEq(int64(2), maxRunTime(2, batteries))

	batteries = []int{10, 10, 3, 5}
	assertEq(int64(8), maxRunTime(3, batteries))
}

// 3 5 10 10
func assertEq(a, b any) {
	if reflect.DeepEqual(a, b) {
		fmt.Printf("Ok: %v\n", a)
	} else {
		fmt.Printf("Failed: %v != %v\n", a, b)
	}
}

func maxRunTime(n int, batteries []int) int64 {
	slices.Sort(batteries)
	var remain int
	for i := range len(batteries) - n {
		remain += batteries[i]
	}

	start := len(batteries) - n
	for j := start + 1; j < len(batteries); j++ {
		need := (batteries[j] - batteries[j-1]) * (j - start)
		if remain < need {
			return int64(batteries[j-1]) + int64(remain/(j-start))
		}
		remain -= need
	}

	return int64(batteries[len(batteries)-1]) + int64(remain/n)
}
