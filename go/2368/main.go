package main

import "fmt"

func main() {
	fmt.Println(reachableNodes(7, [][]int{{0, 1}, {1, 2}, {3, 1}, {4, 0}, {0, 5}, {5, 6}}, []int{4, 5}))
	fmt.Println(reachableNodes(7, [][]int{{0, 1}, {0, 2}, {0, 5}, {0, 4}, {3, 2}, {6, 5}}, []int{4, 2, 1}))
}

func reachableNodes(n int, edges [][]int, restricted []int) int {
	avoided := map[int]bool{}
	for _, n := range restricted {
		avoided[n] = true
	}

	g := map[int][]int{}
	for _, e := range edges {
		g[e[0]] = append(g[e[0]], e[1])
		g[e[1]] = append(g[e[1]], e[0])
	}

	seen := make([]bool, n)
	var dfs func(i int) int
	dfs = func(i int) int {
		var res int
		for _, nei := range g[i] {
			if !seen[nei] && !avoided[nei] {
				seen[nei] = true
				res += dfs(nei)
			}
		}
		return res + 1
	}

	seen[0] = true
	return dfs(0)
}
