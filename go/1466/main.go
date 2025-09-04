package main

import "fmt"

func main() {
	fmt.Println(minReorder(6, [][]int{{0, 1}, {1, 3}, {2, 3}, {4, 0}, {4, 5}}))
	fmt.Println(minReorder(5, [][]int{{1, 0}, {1, 2}, {3, 2}, {3, 4}}))
	fmt.Println(minReorder(3, [][]int{{1, 0}, {2, 0}}))
}

func minReorder(n int, connections [][]int) int {
	g := map[int][]int{}
	m := map[[2]int]bool{}
	for _, c := range connections {
		g[c[0]] = append(g[c[0]], c[1])
		g[c[1]] = append(g[c[1]], c[0])
		m[[2]int{c[0], c[1]}] = true
	}

	var res int
	seen := make([]bool, n)
	var dfs func(i int)
	dfs = func(i int) {
		for _, nei := range g[i] {
			if !seen[nei] {
				seen[nei] = true
				if m[[2]int{i, nei}] {
					res++
				}
				dfs(nei)
			}
		}
	}

	seen[0] = true
	dfs(0)

	return res
}
