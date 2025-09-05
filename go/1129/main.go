package main

import "fmt"

func main() {
	fmt.Println(shortestAlternatingPaths(3, [][]int{{0, 1}, {1, 2}}, [][]int{}))
	fmt.Println(shortestAlternatingPaths(3, [][]int{{0, 1}}, [][]int{{2, 1}}))
}

func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int {
	const (
		RED  = 0
		BLUE = 1
	)

	g := [2][][]int{make([][]int, n), make([][]int, n)}
	for _, e := range redEdges {
		g[RED][e[0]] = append(g[RED][e[0]], e[1])
	}
	for _, e := range blueEdges {
		g[BLUE][e[0]] = append(g[BLUE][e[0]], e[1])
	}

	type node struct {
		n, color int // 0 - red, 1 - blue
	}

	res := make([]int, n)
	for i := range n {
		res[i] = -1
	}
	res[0] = 0

	seen := map[node]bool{} // key 1: node as state
	seen[node{0, 0}] = true
	seen[node{0, 1}] = true
	queue := []node{{0, 0}, {0, 1}} // key 2: red and blue node 0

	step := 1
	for len(queue) > 0 {
		for range len(queue) {
			cur := queue[0]
			queue = queue[1:]
			for _, nei := range g[1-cur.color][cur.n] {
				if !seen[node{nei, 1 - cur.color}] {
					seen[node{nei, 1 - cur.color}] = true
					queue = append(queue, node{nei, 1 - cur.color})
					if res[nei] == -1 { // key 3: keep the smallest (BFS the first is the smallest)
						res[nei] = step
					}
				}
			}
		}
		step++
	}

	return res
}
