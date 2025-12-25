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

func assertEq(a, b any) {
	if reflect.DeepEqual(a, b) {
		fmt.Printf("Ok: %v\n", a)
	} else {
		fmt.Printf("Failed: %v != %v\n", a, b)
	}
}

func maxRunTime(n int, batteries []int) int64 {
	slices.SortFunc(batteries, func(a, b int) int { return b - a })
	live := batteries[:n]
	slices.Reverse(live)

	var extra int
	for i := n; i < len(batteries); i++ {
		extra += batteries[i]
	}

	for i := range n - 1 {
		diff := live[i+1] - live[i]
		if extra < diff*(i+1) {
			return int64(live[i] + extra/(i+1))
		}
		extra -= diff * (i + 1)
	}

	return int64(live[len(live)-1] + extra/n)
}
