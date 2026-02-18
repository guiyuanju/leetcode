package main

import (
	"fmt"
	"reflect"
)

func main() {
	assertEq([]int{4, 2, 4, 2, 3}, finalPrices([]int{8, 4, 6, 2, 3}))
	assertEq([]int{1, 2, 3, 4, 5}, finalPrices([]int{1, 2, 3, 4, 5}))
	assertEq([]int{9, 0, 1, 6}, finalPrices([]int{10, 1, 1, 6}))
}

func assertEq(a, b any) {
	if reflect.DeepEqual(a, b) {
		fmt.Printf("Ok: %v\n", a)
	} else {
		fmt.Printf("Failed: %v != %v\n", a, b)
	}
}

func finalPrices(prices []int) []int {
	res := make([]int, len(prices))
	mono := []int{}
	for i, p := range prices {
		for len(mono) > 0 && p <= prices[mono[len(mono)-1]] {
			idx := mono[len(mono)-1]
			res[idx] = prices[idx] - p
			mono = mono[:len(mono)-1]
		}
		mono = append(mono, i)
	}
	for _, idx := range mono {
		res[idx] = prices[idx]
	}
	return res
}
