package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func main() {
	cells := parse2DArray("[[1,1],[2,1],[1,2],[2,2]]")
	assertEq(2, latestDayToCross(2, 2, cells))

	cells = parse2DArray("[[1,1],[1,2],[2,1],[2,2]]")
	assertEq(1, latestDayToCross(2, 2, cells))

	cells = parse2DArray("[[1,2],[2,1],[3,3],[2,2],[1,1],[1,3],[2,3],[3,2],[3,1]]")
	assertEq(3, latestDayToCross(3, 3, cells))
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

func assertEq(a, b any) {
	if reflect.DeepEqual(a, b) {
		fmt.Printf("Ok: %v\n", a)
	} else {
		fmt.Printf("Failed: %v != %v\n", a, b)
	}
}

func latestDayToCross(row int, col int, cells [][]int) int {
	g := make([][]int, row)
	for i := range row {
		g[i] = make([]int, col)
	}
	for i, c := range cells {
		g[c[0]-1][c[1]-1] = i + 1
	}

	check := func(guess int) bool {
		valid := func(r, c int) bool {
			return 0 <= r && r < row && 0 <= c && c < col
		}
		directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

		queue := [][2]int{}
		seen := map[[2]int]bool{}
		for i := range col {
			if guess < g[0][i] {
				queue = append(queue, [2]int{0, i})
				seen[[2]int{0, i}] = true
			}
		}
		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]
			if cur[0] == row-1 {
				return true
			}
			for _, dir := range directions {
				nr, nc := dir[0]+cur[0], dir[1]+cur[1]
				if valid(nr, nc) && !seen[[2]int{nr, nc}] && guess < g[nr][nc] {
					seen[[2]int{nr, nc}] = true
					queue = append(queue, [2]int{nr, nc})
				}
			}
		}
		return false
	}

	lo := 0
	hi := len(cells) + 1
	for lo < hi {
		mid := lo + (hi-lo)/2
		if check(mid) {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo - 1
}
