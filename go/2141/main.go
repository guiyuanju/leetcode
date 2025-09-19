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
	check := func(g int) bool {
		var res int
		for _, n := range batteries {
			res += min(n, g)
		}
		return res >= n*g
	}

	var lo, hi int
	for _, n := range batteries {
		hi += n
	}
	hi = hi/n + 1

	for lo < hi {
		mid := lo + (hi-lo)/2
		if check(mid) {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return int64(lo - 1)
}
