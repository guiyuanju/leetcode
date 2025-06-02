package main

import "fmt"

func main() {
	roads := [][]int{{0, 1}, {0, 3}, {1, 2}, {1, 3}}
	assertEq(4, maximalNetworkRank(4, roads))

	roads = [][]int{{0, 1}, {0, 3}, {1, 2}, {1, 3}, {2, 3}, {2, 4}}
	assertEq(5, maximalNetworkRank(5, roads))

	roads = [][]int{{0, 1}, {1, 2}, {2, 3}, {2, 4}, {5, 6}, {5, 7}}
	assertEq(5, maximalNetworkRank(8, roads))
}

func assertEq[T comparable](a, b T) {
	if a != b {
		fmt.Printf("Failed: %v != %v\n", a, b)
	} else {
		fmt.Printf("Ok: %v\n", a)
	}
}

func maximalNetworkRank(n int, roads [][]int) int {
	graph := map[int]map[int]bool{}
	for i := range n {
		graph[i] = map[int]bool{}
	}
	for _, r := range roads {
		graph[r[0]][r[1]] = true
		graph[r[1]][r[0]] = true
	}

	maxRank := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			rank := len(graph[i]) + len(graph[j])
			if graph[i][j] {
				rank--
			}
			maxRank = max(maxRank, rank)
		}
	}

	return maxRank
}
