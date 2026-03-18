package main

import (
	"fmt"
	"reflect"
)

func main() {
	assertEq(4, longestIncreasingPath([][]int{{9, 9, 4}, {6, 6, 8}, {2, 1, 1}}))
	assertEq(4, longestIncreasingPath([][]int{{3, 4, 5}, {3, 2, 6}, {2, 2, 1}}))
	assertEq(1, longestIncreasingPath([][]int{{1}}))
	assertEq(4, longestIncreasingPath([][]int{{7, 7, 5}, {2, 4, 6}, {8, 2, 0}}))
	assertEq(3, longestIncreasingPath([][]int{{0}, {1}, {5}, {5}}))
}

func assertEq(a, b any) {
	if reflect.DeepEqual(a, b) {
		fmt.Printf("Ok: %v\n", a)
	} else {
		fmt.Printf("Failed: %v != %v\n", a, b)
	}
}

func longestIncreasingPath(matrix [][]int) int {
	m := len(matrix)
	n := len(matrix[0])

	valid := func(r, c int) bool {
		return 0 <= r && r < m && 0 <= c && c < n
	}
	directions := [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}

	indegree := map[[2]int]int{}
	g := map[[2]int][][2]int{}
	for i := range m {
		for j := range n {
			for _, dir := range directions {
				nr, nc := i+dir[0], j+dir[1]
				if valid(nr, nc) && matrix[nr][nc] > matrix[i][j] {
					g[[2]int{i, j}] = append(g[[2]int{i, j}], [2]int{nr, nc})
					indegree[[2]int{nr, nc}]++
				}
			}
		}
	}

	q := [][2]int{}
	for i := range m {
		for j := range n {
			if indegree[[2]int{i, j}] == 0 {
				q = append(q, [2]int{i, j})
			}
		}
	}

	length := 0
	for len(q) > 0 {
		for range q {
			cur := q[0]
			q = q[1:]
			for _, nei := range g[cur] {
				indegree[nei]--
				if indegree[nei] == 0 {
					q = append(q, nei)
				}
			}
		}
		length++
	}

	return length
}
