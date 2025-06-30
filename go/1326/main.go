package main

import (
	"fmt"
	"slices"
)

func main() {
	n := 5
	ranges := []int{3, 4, 1, 1, 0, 0}
	assertEq(1, minTaps(n, ranges))

	n = 3
	ranges = []int{0, 0, 0, 0}
	assertEq(-1, minTaps(n, ranges))

	n = 7
	ranges = []int{1, 2, 1, 0, 2, 1, 0, 1}
	assertEq(3, minTaps(n, ranges))

	n = 8
	ranges = []int{4, 0, 0, 0, 0, 0, 0, 0, 4}
	assertEq(2, minTaps(n, ranges))
}

func assertEq[T comparable](a, b T) {
	if a != b {
		fmt.Printf("Failed: %v != %v\n", a, b)
	} else {
		fmt.Printf("Ok: %v\n", a)
	}
}

func minTaps(n int, ranges []int) int {
	realRanges := make([][]int, len(ranges))
	for i, r := range ranges {
		realRanges[i] = []int{i - r, i + r}
	}
	slices.SortFunc(realRanges, func(a, b []int) int { return a[0] - b[0] })

	var i, end, res int
	for end < n {
		var tmp int
		for ; i < len(realRanges); i++ {
			if realRanges[i][0] > end {
				break
			}
			tmp = max(tmp, realRanges[i][1])
		}
		if tmp == end {
			return -1
		}
		end = tmp
		res++
	}
	return res
}
