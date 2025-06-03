package main

import "fmt"

func main() {
	n := 8
	edgeList := [][]int{{0, 3}, {0, 4}, {1, 3}, {2, 4}, {2, 7}, {3, 5}, {3, 6}, {3, 7}, {4, 6}}
	fmt.Println(getAncestors(n, edgeList))

	n = 5
	edgeList = [][]int{{0, 1}, {0, 2}, {0, 3}, {0, 4}, {1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}}
	fmt.Println(getAncestors(n, edgeList))
}

func getAncestors(n int, edges [][]int) [][]int {
	graph := map[int][]int{}
	for _, e := range edges {
		graph[e[0]] = append(graph[e[0]], e[1])
	}

	res := make([][]int, n)
	seen := make([]bool, n)
	var dfs func(n, root int)
	dfs = func(n int, root int) {
		for _, nei := range graph[n] {
			if !seen[nei] {
				seen[nei] = true
				res[nei] = append(res[nei], root)
				dfs(nei, root)
			}
		}
	}

	for i := range n {
		dfs(i, i)
		for i := range seen {
			seen[i] = false
		}
	}

	return res
}
