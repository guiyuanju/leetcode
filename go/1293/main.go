package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func main() {
	assertEq(6, shortestPath(makeGrid("[[0,0,0],[1,1,0],[0,0,0],[0,1,1],[0,0,0]]"), 1))
	assertEq(-1, shortestPath(makeGrid("[[0,1,1],[1,1,1],[1,0,0]]"), 1))
	assertEq(20, shortestPath(makeGrid("[[0,0,0,0,0,0,0,0,0,0],[0,1,1,1,1,1,1,1,1,0],[0,1,0,0,0,0,0,0,0,0],[0,1,0,1,1,1,1,1,1,1],[0,1,0,0,0,0,0,0,0,0],[0,1,1,1,1,1,1,1,1,0],[0,1,0,0,0,0,0,0,0,0],[0,1,0,1,1,1,1,1,1,1],[0,1,0,1,1,1,1,0,0,0],[0,1,0,0,0,0,0,0,1,0],[0,1,1,1,1,1,1,0,1,0],[0,0,0,0,0,0,0,0,1,0]]"), 1))
}

func assertEq(a, b any) {
	if reflect.DeepEqual(a, b) {
		fmt.Printf("Ok: %v\n", a)
	} else {
		fmt.Printf("Failed: %v != %v\n", a, b)
	}
}

func shortestPath(grid [][]int, k int) int {
	m := len(grid)
	n := len(grid[0])
	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	valid := func(r, c int) bool {
		return 0 <= r && r < m && 0 <= c && c < n
	}

	type node struct {
		r, c, left, step int
	}

	seen := map[[3]int]bool{}
	seen[[3]int{0, 0, k}] = true
	queue := []node{{0, 0, k, 0}}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if cur.r == m-1 && cur.c == n-1 {
			return cur.step
		}
		for _, nei := range directions {
			nextRow, nextCol := cur.r+nei[0], cur.c+nei[1]
			if valid(nextRow, nextCol) {
				if grid[nextRow][nextCol] == 1 && cur.left > 0 && !seen[[3]int{nextRow, nextCol, cur.left - 1}] {
					seen[[3]int{nextRow, nextCol, cur.left - 1}] = true
					queue = append(queue, node{nextRow, nextCol, cur.left - 1, cur.step + 1})
				} else if grid[nextRow][nextCol] == 0 && !seen[[3]int{nextRow, nextCol, cur.left}] {
					seen[[3]int{nextRow, nextCol, cur.left}] = true
					queue = append(queue, node{nextRow, nextCol, cur.left, cur.step + 1})
				}
			}
		}
	}

	return -1
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
