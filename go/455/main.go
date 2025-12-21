package main

import (
	"fmt"
	"slices"
)

func main() {
	g := []int{1, 2, 3}
	s := []int{1, 1}
	assertEq(1, findContentChildren(g, s))

	g = []int{1, 2}
	s = []int{1, 2, 3}
	assertEq(2, findContentChildren(g, s))

	g = []int{10, 9, 8, 7}
	s = []int{5, 6, 7, 8}
	assertEq(2, findContentChildren(g, s))
}

func assertEq[T comparable](a, b T) {
	if a != b {
		fmt.Printf("Failed: %v != %v\n", a, b)
	} else {
		fmt.Printf("Ok: %v\n", a)
	}
}

func findContentChildren(g []int, s []int) int {
	slices.Sort(g)
	slices.Sort(s)

	var res, i, j int
	for i < len(g) && j < len(s) {
		if g[i] <= s[j] {
			res++
			i++
		}
		j++
	}

	return res
}
