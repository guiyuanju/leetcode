package main

import (
	"fmt"
	"slices"
)

func main() {
	weight := []int{100, 200, 150, 1000}
	assertEq(4, maxNumberOfApples(weight))

	weight = []int{900, 950, 800, 1000, 700, 800}
	assertEq(5, maxNumberOfApples(weight))
}

func assertEq[T comparable](a, b T) {
	if a != b {
		fmt.Printf("Failed: %v != %v\n", a, b)
	} else {
		fmt.Printf("Ok: %v\n", a)
	}
}

func maxNumberOfApples(weight []int) int {
	slices.Sort(weight)
	var res, sum int
	for _, w := range weight {
		if sum+w > 5000 {
			return res
		}
		sum += w
		res++
	}
	return res
}
