package main

import "fmt"

func main() {
	fmt.Println(longestIncreasingPath([][]int{{9, 9, 4}, {6, 6, 8}, {2, 1, 1}}))
	fmt.Println(longestIncreasingPath([][]int{{3, 4, 5}, {3, 2, 6}, {2, 2, 1}}))
	fmt.Println(longestIncreasingPath([][]int{{1}}))
}

func longestIncreasingPathDP(matrix [][]int) int {
	m := len(matrix)
	n := len(matrix[0])

	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	valid := func(r, c int) bool {
		return 0 <= r && r < m && 0 <= c && c < n
	}

	memo := map[[2]int]int{}
	var dp func(i, j int) int
	dp = func(i, j int) int {
		if v, ok := memo[[2]int{i, j}]; ok {
			return v
		}

		var res int
		for _, dir := range directions {
			nextRow, nextCol := i+dir[0], j+dir[1]
			if valid(nextRow, nextCol) && matrix[nextRow][nextCol] > matrix[i][j] {
				res = max(res, dp(nextRow, nextCol))
			}
		}

		memo[[2]int{i, j}] = res + 1

		return res + 1
	}

	var res int
	for i := range m {
		for j := range n {
			res = max(res, dp(i, j))
		}
	}
	return res
}

// topological sort
func longestIncreasingPath(matrix [][]int) int {
	m := len(matrix)
	n := len(matrix[0])

	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	valid := func(r, c int) bool {
		return 0 <= r && r < m && 0 <= c && c < n
	}

	ind := make([][]int, m)
	for i := range ind {
		ind[i] = make([]int, n)
	}

	q := [][3]int{}
	g := make(map[[2]int][][2]int, m*n)
	for i := range m {
		for j := range n {
			var degree int
			for _, dir := range directions {
				nextRow, nextCol := i+dir[0], j+dir[1]
				if valid(nextRow, nextCol) && matrix[i][j] > matrix[nextRow][nextCol] {
					g[[2]int{nextRow, nextCol}] = append(g[[2]int{nextRow, nextCol}], [2]int{i, j})
					degree++
				}
			}
			ind[i][j] = degree
			if degree == 0 {
				q = append(q, [3]int{i, j, 1})
			}
		}
	}

	var res int
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		res = max(res, cur[2])
		for _, nei := range g[[2]int{cur[0], cur[1]}] {
			ind[nei[0]][nei[1]]--
			if ind[nei[0]][nei[1]] == 0 {
				q = append(q, [3]int{nei[0], nei[1], cur[2] + 1})
			}
		}
	}

	return res
}
