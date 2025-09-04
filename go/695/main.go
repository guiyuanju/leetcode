package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(maxAreaOfIsland(makeGrid("[[0,0,1,0,0,0,0,1,0,0,0,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,1,1,0,1,0,0,0,0,0,0,0,0],[0,1,0,0,1,1,0,0,1,0,1,0,0],[0,1,0,0,1,1,0,0,1,1,1,0,0],[0,0,0,0,0,0,0,0,0,0,1,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,0,0,0,0,0,0,1,1,0,0,0,0]]")))
	fmt.Println(maxAreaOfIsland(makeGrid("[[0,0,0,0,0,0,0,0]]")))
}

func maxAreaOfIsland(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	valid := func(r, c int) bool {
		return 0 <= r && r < m && 0 <= c && c < n && grid[r][c] == 1
	}

	var dfs func(r, w int) int
	dfs = func(r, w int) int {
		var res int
		for _, nei := range directions {
			nextRow, nextCol := r+nei[0], w+nei[1]
			if valid(nextRow, nextCol) {
				grid[nextRow][nextCol] = 0
				res += dfs(nextRow, nextCol)
			}
		}
		return res + 1
	}

	var res int
	for i := range m {
		for j := range n {
			if grid[i][j] == 1 {
				grid[i][j] = 0
				res = max(res, dfs(i, j))
			}
		}
	}

	return res
}

func printGrid(g [][]int) {
	for _, r := range g {
		for _, v := range r {
			fmt.Printf("%6d", v)
		}
		fmt.Println()
	}
}

func makeGrid(s string) [][]int {
	s = s[1 : len(s)-1]
	parts := strings.Split(s, "],[")
	var res [][]int
	for _, p := range parts {
		p = strings.Trim(p, "[]")
		if len(p) == 0 {
			res = append(res, []int{})
			continue
		}
		ints := strings.Split(p, ",")
		var cur []int
		for _, i := range ints {
			n, err := strconv.ParseInt(i, 10, 0)
			if err != nil {
				panic(err)
			}
			cur = append(cur, int(n))
		}
		res = append(res, cur)
	}
	return res
}
