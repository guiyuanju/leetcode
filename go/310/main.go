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

	graph := make(map[int][]int, n)
	degrees := make([]int, n)
	for _, e := range edges {
		graph[e[0]] = append(graph[e[0]], e[1])
		graph[e[1]] = append(graph[e[1]], e[0])
		degrees[e[0]]++
		degrees[e[1]]++
	}

	queue := []int{}
	for i, d := range degrees {
		if d == 1 {
			queue = append(queue, i)
		}
	}

	remain := n - len(queue)
	for remain > 0 {
		for range len(queue) {
			cur := queue[0]
			queue = queue[1:]
			for _, nei := range graph[cur] {
				degrees[nei]--
				if degrees[nei] == 1 {
					queue = append(queue, nei)
					remain--
				}
			}
		}
	}

	return queue
}
