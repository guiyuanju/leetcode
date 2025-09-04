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
	type node struct {
		r, c, s int
	}
	m := len(mat)
	n := len(mat[0])
	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	valid := func(r, c int) bool {
		return 0 <= r && r < m && 0 <= c && c < n
	}

	seen := make([][]bool, m)
	for i := range m {
		seen[i] = make([]bool, n)
	}

	queue := []node{}
	for i := range m {
		for j := range n {
			if mat[i][j] == 0 {
				queue = append(queue, node{i, j, 0})
				seen[i][j] = true
			}
		}
	}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		for _, nei := range directions {
			nextRow, nextCol := cur.r+nei[0], cur.c+nei[1]
			if valid(nextRow, nextCol) && !seen[nextRow][nextCol] {
				seen[nextRow][nextCol] = true
				queue = append(queue, node{nextRow, nextCol, cur.s + 1})
				mat[nextRow][nextCol] = cur.s + 1
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
