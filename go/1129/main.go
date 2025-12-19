package main

import (
	"fmt"
	"reflect"
)

func main() {
	assertEq([]int{0, 1, -1}, shortestAlternatingPaths(3, [][]int{{0, 1}, {1, 2}}, [][]int{}))
	assertEq([]int{0, 1, -1}, shortestAlternatingPaths(3, [][]int{{0, 1}}, [][]int{{2, 1}}))
	assertEq([]int{0, 1, 2}, shortestAlternatingPaths(3, [][]int{{0, 1}}, [][]int{{1, 2}}))
}

func assertEq(a, b any) {
	if reflect.DeepEqual(a, b) {
		fmt.Printf("Ok: %v\n", a)
	} else {
		fmt.Printf("Failed: %v != %v\n", a, b)
	}
}

func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int {
	g := [2]map[int][]int{{}, {}}
	for _, e := range redEdges {
		g[0][e[0]] = append(g[0][e[0]], e[1])
	}
	for _, e := range blueEdges {
		g[1][e[0]] = append(g[1][e[0]], e[1])
	}

	type node struct {
		n, c, s int
	}

	seen := map[[2]int]bool{}
	seen[[2]int{0, 0}] = true
	seen[[2]int{0, 1}] = true
	q := []node{{0, 0, 0}, {0, 1, 0}}

	res := make([]int, n)
	for i := range res {
		res[i] = -1
	}

	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		if res[cur.n] == -1 {
			res[cur.n] = cur.s
		}
		for _, nei := range g[1-cur.c][cur.n] {
			if !seen[[2]int{nei, 1 - cur.c}] {
				seen[[2]int{nei, 1 - cur.c}] = true
				q = append(q, node{nei, 1 - cur.c, cur.s + 1})
			}
		}
	}

	return res
}
