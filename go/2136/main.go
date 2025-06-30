package main

import (
	"fmt"
	"slices"
)

func main() {
	plantTime := []int{1, 4, 3}
	grwoTime := []int{2, 3, 1}
	assertEq(9, earliestFullBloom(plantTime, grwoTime))

	plantTime = []int{1, 2, 3, 2}
	grwoTime = []int{2, 1, 2, 1}
	assertEq(9, earliestFullBloom(plantTime, grwoTime))

	plantTime = []int{1}
	grwoTime = []int{1}
	assertEq(2, earliestFullBloom(plantTime, grwoTime))
}

func assertEq[T comparable](a, b T) {
	if a != b {
		fmt.Printf("Failed: %v != %v\n", a, b)
	} else {
		fmt.Printf("Ok: %v\n", a)
	}
}

func earliestFullBloom(plantTime []int, growTime []int) int {
	indices := make([]int, len(plantTime))
	for i := range indices {
		indices[i] = i
	}
	slices.SortFunc(indices, func(a, b int) int { return -(growTime[a] - growTime[b]) })

	var res, cur int
	for _, i := range indices {
		res = max(res, cur+plantTime[i]+growTime[i])
		cur += plantTime[i]
	}
	return res
}
