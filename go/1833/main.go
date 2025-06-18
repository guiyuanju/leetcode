package main

import (
	"fmt"
	"slices"
)

func main() {
	costs := []int{1, 3, 2, 4, 1}
	coins := 7
	assertEq(4, maxIceCream(costs, coins))

	costs = []int{10, 6, 8, 7, 7, 8}
	coins = 5
	assertEq(0, maxIceCream(costs, coins))

	costs = []int{1, 6, 3, 1, 2, 5}
	coins = 20
	assertEq(6, maxIceCream(costs, coins))
}

func assertEq[T comparable](a, b T) {
	if a != b {
		fmt.Printf("Failed: %v != %v\n", a, b)
	} else {
		fmt.Printf("Ok: %v\n", a)
	}
}

func maxIceCream(costs []int, coins int) int {
	slices.Sort(costs)

	var res int
	for _, c := range costs {
		if coins < c {
			return res
		}
		res++
		coins -= c
	}

	return res
}
