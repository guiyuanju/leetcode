package main

import (
	"fmt"
	"reflect"
)

func main() {
	assertEq([]int{2, 6}, mostCompetitive([]int{3, 5, 2, 6}, 2))
	assertEq([]int{2, 3, 3, 4}, mostCompetitive([]int{2, 4, 3, 3, 5, 4, 9, 6}, 4))
	assertEq([]int{3, 1, 0}, mostCompetitive([]int{3, 4, 5, 1, 0}, 3))
}

func assertEq(a, b any) {
	if reflect.DeepEqual(a, b) {
		fmt.Printf("Ok: %v\n", a)
	} else {
		fmt.Printf("Failed: %v != %v\n", a, b)
	}
}

func mostCompetitive(nums []int, k int) []int {
	var res []int
	for i, n := range nums {
		for len(res) > 0 && n < res[len(res)-1] && len(res)+len(nums)-i-1 >= k {
			res = res[:len(res)-1]
		}
		res = append(res, n)
	}
	return res[:k]
}

// [3 4 5 1 0] 3 1 0
