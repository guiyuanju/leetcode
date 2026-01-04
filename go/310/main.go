package main

import (
	"fmt"
	"reflect"
)

func main() {
	assertEq([]int{1}, findMinHeightTrees(4, [][]int{{1, 0}, {1, 2}, {1, 3}}))
	assertEq([]int{3, 4}, findMinHeightTrees(6, [][]int{{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4}}))
	assertEq([]int{0}, findMinHeightTrees(1, [][]int{}))
}

func assertEq(a, b any) {
	if reflect.DeepEqual(a, b) {
		fmt.Printf("Ok: %v\n", a)
	} else {
		fmt.Printf("Failed: %v != %v\n", a, b)
	}
}

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}

	g := map[int][]int{}
	deg := make([]int, n)
	for _, e := range edges {
		deg[e[0]]++
		deg[e[1]]++
		g[e[0]] = append(g[e[0]], e[1])
		g[e[1]] = append(g[e[1]], e[0])
	}

	q := []int{}
	for n, d := range deg {
		if d == 1 {
			q = append(q, n)
		}
	}

	left := n - len(q)
	for left > 0 {
		for range len(q) {
			cur := q[0]
			q = q[1:]
			for _, nei := range g[cur] {
				deg[nei]--
				if deg[nei] == 1 {
					left--
					q = append(q, nei)
				}
			}
		}
	}

	return q
}
