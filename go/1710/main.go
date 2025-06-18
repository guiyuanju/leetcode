package main

import (
	"fmt"
	"slices"
)

func main() {
	boxType := [][]int{{1, 3}, {2, 2}, {3, 1}}
	truckSize := 4
	assertEq(8, maximumUnits(boxType, truckSize))

	boxType = [][]int{{5, 10}, {2, 5}, {4, 7}, {3, 9}}
	truckSize = 10
	assertEq(91, maximumUnits(boxType, truckSize))
}

func assertEq[T comparable](a, b T) {
	if a != b {
		fmt.Printf("Failed: %v != %v\n", a, b)
	} else {
		fmt.Printf("Ok: %v\n", a)
	}
}

func maximumUnits(boxTypes [][]int, truckSize int) int {
	slices.SortFunc(boxTypes, func(a, b []int) int {
		return b[1] - a[1]
	})

	var res int
	var i int
	for truckSize > 0 && i < len(boxTypes) {
		if boxTypes[i][0] >= truckSize {
			res += boxTypes[i][1] * truckSize
			break
		}
		res += boxTypes[i][1] * boxTypes[i][0]
		truckSize -= boxTypes[i][0]
		i++
	}

	return res
}
