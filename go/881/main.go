package main

import (
	"fmt"
	"slices"
)

func main() {
	people := []int{1, 2}
	limit := 3
	assertEq(1, numRescueBoats(people, limit))

	people = []int{3, 2, 2, 1}
	limit = 3
	assertEq(3, numRescueBoats(people, limit))

	people = []int{3, 5, 3, 4}
	limit = 5
	assertEq(4, numRescueBoats(people, limit))

	people = []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 4, 4}
	limit = 4
	assertEq(8, numRescueBoats(people, limit))
}

func assertEq[T comparable](a, b T) {
	if a != b {
		fmt.Printf("Failed: %v != %v\n", a, b)
	} else {
		fmt.Printf("Ok: %v\n", a)
	}
}

func numRescueBoats(people []int, limit int) int {
	slices.Sort(people)
	slices.Reverse(people)
	i := 0
	j := len(people) - 1
	var res int
	for i <= j {
		if people[i]+people[j] <= limit {
			j--
		}
		i++
		res++
	}
	return res
}
