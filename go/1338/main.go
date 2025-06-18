package main

import (
	"fmt"
	"slices"
)

func main() {
	arr := []int{3, 3, 3, 3, 5, 5, 5, 2, 2, 7}
	assertEq(2, minSetSize(arr))

	arr = []int{7, 7, 7, 7, 7, 7}
	assertEq(1, minSetSize(arr))
}

func assertEq[T comparable](a, b T) {
	if a != b {
		fmt.Printf("Failed: %v != %v\n", a, b)
	} else {
		fmt.Printf("Ok: %v\n", a)
	}
}

func minSetSize(arr []int) int {
	counts := map[int]int{}
	for _, v := range arr {
		counts[v]++
	}

	var countArr []int
	for _, c := range counts {
		countArr = append(countArr, c)
	}
	slices.SortFunc(countArr, func(a, b int) int { return b - a })

	var count, sum int
	for _, c := range countArr {
		sum += c
		count++
		if sum >= len(arr)/2 {
			return count
		}
	}
	return count
}
