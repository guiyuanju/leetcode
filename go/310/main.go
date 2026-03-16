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
	degrees := make([]int, n)
	g := map[int][]int{}
	for _, e := range edges {
		g[e[0]] = append(g[e[0]], e[1])
		g[e[1]] = append(g[e[1]], e[0])
		degrees[e[0]]++
		degrees[e[1]]++
	}

	q := []int{}
	for i, d := range degrees {
		if d <= 1 {
			q = append(q, i)
		}
	}

	var res []int
	for len(q) > 0 {
		res = q
		for range q {
			cur := q[0]
			q = q[1:]
			degrees[cur]--
			for _, nei := range g[cur] {
				degrees[nei]--
				if degrees[nei] == 1 {
					q = append(q, nei)
				}
			}
		}
	}

	return res
}
