package main

import "fmt"

func main() {
	fmt.Println(longestIncreasingPath([][]int{{9, 9, 4}, {6, 6, 8}, {2, 1, 1}}))
	fmt.Println(longestIncreasingPath([][]int{{3, 4, 5}, {3, 2, 6}, {2, 2, 1}}))
	fmt.Println(longestIncreasingPath([][]int{{1}}))
}

func longestIncreasingPath(matrix [][]int) int {
	m := len(matrix)
	n := len(matrix[0])
	valid := func(r, c int) bool {
		return 0 <= r && r < m && 0 <= c && c < n
	}
	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	// compute indgrees
	indegrees := make([][]int, m)
	for i := range indegrees {
		indegrees[i] = make([]int, n)
	}
	for r := range m {
		for c := range n {
			for _, dir := range directions {
				nr, nc := r+dir[0], c+dir[1]
				if valid(nr, nc) && matrix[r][c] < matrix[nr][nc] {
					indegrees[nr][nc]++
				}
			}
		}
	}

	// find the longest path
	queue := [][2]int{}
	for r := range m {
		for c := range n {
			if indegrees[r][c] == 0 {
				queue = append(queue, [2]int{r, c})
			}
		}
	}
	var res int
	for len(queue) > 0 {
		for range len(queue) {
			cur := queue[0]
			queue = queue[1:]
			for _, dir := range directions {
				nr, nc := cur[0]+dir[0], cur[1]+dir[1]
				if valid(nr, nc) && matrix[cur[0]][cur[1]] < matrix[nr][nc] {
					indegrees[nr][nc]--
					if indegrees[nr][nc] == 0 {
						queue = append(queue, [2]int{nr, nc})
					}
				}
			}
		}
		res++
	}

	return res
}
