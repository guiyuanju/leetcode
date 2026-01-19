package main

import (
	"fmt"
	"reflect"
)

func main() {
	assertEq(49, maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	assertEq(1, maxArea([]int{1, 1}))
}

func maxArea(height []int) int {
	var res int
	i := 0
	j := len(height) - 1
	for i < j {
		res = max(res, (j-i)*min(height[i], height[j]))
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}

	return res
}

func assertEq(a, b any) {
	if reflect.DeepEqual(a, b) {
		fmt.Printf("Ok: %v\n", a)
	} else {
		fmt.Printf("Failed: %v != %v\n", a, b)
	}
}
