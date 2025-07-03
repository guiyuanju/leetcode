package main

import (
	"fmt"
	"math"
	"slices"
)

func main() {
	piles := []int{3, 6, 7, 11}
	h := 8
	assertEq(4, minEatingSpeed(piles, h))

	piles = []int{30, 11, 23, 4, 20}
	h = 5
	assertEq(30, minEatingSpeed(piles, h))

	piles = []int{30, 11, 23, 4, 20}
	h = 6
	assertEq(23, minEatingSpeed(piles, h))
}

func assertEq[T comparable](a, b T) {
	if a != b {
		fmt.Printf("Failed: %v != %v\n", a, b)
	} else {
		fmt.Printf("Ok: %v\n", a)
	}
}

func minEatingSpeed(piles []int, h int) int {
	check := func(k int) bool {
		hours := 0
		for _, p := range piles {
			hours += int(math.Ceil(float64(p) / float64(k)))
		}
		return hours <= h
	}

	left := 1
	right := slices.Max(piles)
	for left < right {
		mid := left + (right-left)/2
		if check(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
