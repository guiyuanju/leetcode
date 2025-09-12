package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	heights := parse2DArray("[[1,2,2],[3,8,2],[5,3,5]]")
	assertEq(2, minimumEffortPath(heights))

	heights = parse2DArray("[[1,2,3],[3,8,4],[5,3,5]]")
	assertEq(1, minimumEffortPath(heights))

	heights = parse2DArray("[[1,2,1,1,1],[1,2,1,2,1],[1,2,1,2,1],[1,2,1,2,1],[1,1,1,2,1]]")
	assertEq(0, minimumEffortPath(heights))
}

func assertEq[T comparable](a, b T) {
	if a != b {
		fmt.Printf("Failed: %v != %v\n", a, b)
	} else {
		fmt.Printf("Ok: %v\n", a)
	}
}

func parse2DArray(s string) [][]int {
	s = s[1 : len(s)-1]
	parts := strings.Split(s, "],[")
	var res [][]int
	for _, p := range parts {
		p = strings.Trim(p, "[]")
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

// bianry search
func minimumEffortPath(heights [][]int) int {
	m := len(heights)
	n := len(heights[0])

	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	valid := func(r, c int) bool {
		return 0 <= r && r < m && 0 <= c && c < n
	}

	seen := make([][]bool, m)
	for i := range m {
		seen[i] = make([]bool, n)
	}

	var dfs func(r, c, limit int) bool
	dfs = func(r, c, limit int) bool {
		if r == m-1 && c == n-1 {
			return true
		}
		for _, dir := range directions {
			nr, nc := r+dir[0], c+dir[1]
			if valid(nr, nc) && !seen[nr][nc] && abs(heights[nr][nc]-heights[r][c]) <= limit {
				seen[nr][nc] = true
				if dfs(nr, nc, limit) {
					return true
				}
			}
		}
		return false
	}

	// binary search
	var lo, hi int
	for i := range m {
		for j := range n {
			lo = min(lo, heights[i][j])
			hi = max(hi, heights[i][j])
		}
	}

	i := 0
	j := (hi - lo) + 1
	for i < j {
		mid := i + (j-i)/2
		for k := range seen {
			clear(seen[k])
		}
		seen[0][0] = true
		if dfs(0, 0, mid) {
			j = mid
		} else {
			i = mid + 1
		}
	}

	return i
}

func dbg[T any](x T) T {
	fmt.Println("dbg:", x)
	return x
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
