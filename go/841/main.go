package main

import "fmt"

func main() {
	fmt.Println(canVisitAllRooms([][]int{{1}, {2}, {3}, {}}))
	fmt.Println(canVisitAllRooms([][]int{{1, 3}, {3, 0, 1}, {2}, {0}}))
}

func canVisitAllRooms(rooms [][]int) bool {
	seen := make([]bool, len(rooms))
	var dfs func(i int)
	dfs = func(i int) {
		for _, nei := range rooms[i] {
			if !seen[nei] {
				seen[nei] = true
				dfs(nei)
			}
		}
	}

	seen[0] = true
	dfs(0)

	for _, s := range seen {
		if !s {
			return false
		}
	}

	return true
}
