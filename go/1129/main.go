package main

import (
	"fmt"
	"reflect"
)

func main() {
	assertEq([]int{0, 1, -1}, shortestAlternatingPaths(3, [][]int{{0, 1}, {1, 2}}, [][]int{}))
	assertEq([]int{0, 1, -1}, shortestAlternatingPaths(3, [][]int{{0, 1}}, [][]int{{2, 1}}))
	assertEq([]int{0, 1, 2}, shortestAlternatingPaths(3, [][]int{{0, 1}}, [][]int{{1, 2}}))
	assertEq([]int{0, 1, 2, 3, 7}, shortestAlternatingPaths(5, [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}}, [][]int{{1, 2}, {2, 3}, {3, 1}}))
}

func assertEq(a, b any) {
	if reflect.DeepEqual(a, b) {
		fmt.Printf("Ok: %v\n", a)
	} else {
		fmt.Printf("Failed: %v != %v\n", a, b)
	}
}

func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int {
	rg := map[int][]int{}
	bg := map[int][]int{}
	for _, e := range redEdges {
		rg[e[0]] = append(rg[e[0]], e[1])
	}
	for _, e := range blueEdges {
		bg[e[0]] = append(bg[e[0]], e[1])
	}

	type step struct {
		node, step, color int
	}

	res := make([]int, n)
	for i := range res {
		res[i] = -1
	}
	res[0] = 0

	seen := map[[2]int]bool{}
	seen[[2]int{0, 0}] = true
	seen[[2]int{0, 1}] = true

	q := []step{{0, 0, 0}, {0, 0, 1}}
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		g := rg
		if cur.color == 0 {
			g = bg
		}
		if res[cur.node] == -1 {
			res[cur.node] = cur.step
		}
		for _, nei := range g[cur.node] {
			if !seen[[2]int{nei, 1 - cur.color}] {
				seen[[2]int{nei, 1 - cur.color}] = true
				q = append(q, step{nei, cur.step + 1, 1 - cur.color})
			}
		}
	}

	return res
}
