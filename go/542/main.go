package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	printGrid(updateMatrix(makeGrid("[[0,0,0],[0,1,0],[0,0,0]]")))
	printGrid(updateMatrix(makeGrid("[[0,0,0],[0,1,0],[1,1,1]]")))
}

func updateMatrix(mat [][]int) [][]int {
	m := len(mat)
	n := len(mat[0])
	seen := make([][]bool, m)
	for i := range m {
		seen[i] = make([]bool, n)
	}

	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	valid := func(r, c int) bool {
		return 0 <= r && r < m && 0 <= c && c < n
	}

	q := [][2]int{}
	for i := range m {
		for j := range n {
			if mat[i][j] == 0 {
				q = append(q, [2]int{i, j})
				seen[i][j] = true
			}
		}
	}

	var step int
	for len(q) > 0 {
		step++
		for range len(q) {
			cur := q[0]
			q = q[1:]
			for _, dir := range dirs {
				nr, nc := cur[0]+dir[0], cur[1]+dir[1]
				if valid(nr, nc) && !seen[nr][nc] {
					seen[nr][nc] = true
					mat[nr][nc] = step
					q = append(q, [2]int{nr, nc})
				}
			}
		}
	}

	return mat
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
