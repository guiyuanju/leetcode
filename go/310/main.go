package main

import (
	"fmt"
)

func main() {
	fmt.Println(findMinHeightTrees(4, [][]int{{1, 0}, {1, 2}, {1, 3}}))
	fmt.Println(findMinHeightTrees(6, [][]int{{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4}}))
	fmt.Println(findMinHeightTrees(1, [][]int{}))
}

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}

	g := make(map[int][]int, n)
	degree := make([]int, n)
	for _, e := range edges {
		g[e[0]] = append(g[e[0]], e[1])
		g[e[1]] = append(g[e[1]], e[0])
		degree[e[0]]++
		degree[e[1]]++
	}

	q := []int{}
	for n, d := range degree {
		if d == 1 {
			q = append(q, n)
		}
	}

	remain := n - len(q)

	for remain > 0 {
		length := len(q)
		for range length {
			cur := q[0]
			q = q[1:]
			for _, nei := range g[cur] {
				degree[nei]--
				if degree[nei] == 1 {
					q = append(q, nei)
					remain--
				}
			}
		}
	}

	return q
}
