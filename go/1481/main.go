package main

import (
	"fmt"
	"slices"
)

func main() {
	arr := []int{5, 5, 4}
	k := 1
	assertEq(1, findLeastNumOfUniqueInts(arr, k))
}

func assertEq[T comparable](a, b T) {
	if a != b {
		fmt.Printf("Failed: %v != %v\n", a, b)
	} else {
		fmt.Printf("Ok: %v\n", a)
	}
}

func findLeastNumOfUniqueInts(arr []int, k int) int {
	counts := map[int]int{}
	for _, n := range arr {
		counts[n]++
	}

	countsArr := []int{}
	for _, c := range counts {
		countsArr = append(countsArr, c)
	}

	slices.SortFunc(countsArr, func(a, b int) int { return a - b })
	var sum int
	for i, c := range countsArr {
		sum += c
		if sum > k {
			return len(countsArr) - i
		}
		if sum == k {
			return len(countsArr) - i - 1
		}
	}

	return 0
}
