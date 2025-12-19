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
	type node struct {
		r, c, k, s int
	}
	m := len(grid)
	n := len(grid[0])

	dirs := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	valid := func(r, c int) bool {
		return 0 <= r && r < m && 0 <= c && c < n
	}
	status := func(n node) [3]int {
		return [3]int{n.r, n.c, n.k}
	}

	q := []node{{0, 0, k, 0}}
	seen := map[[3]int]bool{}
	seen[[3]int{0, 0, k}] = true
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		if cur.r == m-1 && cur.c == n-1 {
			return cur.s
		}
		for _, dir := range dirs {
			nr, nc := cur.r+dir[0], cur.c+dir[1]
			if !valid(nr, nc) {
				continue
			}
			nn := node{nr, nc, cur.k, cur.s + 1}
			if grid[nr][nc] == 1 {
				if nn.k == 0 {
					continue
				}
				nn.k--
			}
			if !seen[status(nn)] {
				seen[status(nn)] = true
				q = append(q, nn)
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
