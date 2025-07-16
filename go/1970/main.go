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
	water := make([][]int, row)
	for i := range water {
		water[i] = make([]int, col)
	}
	for i, c := range cells {
		water[c[0]-1][c[1]-1] = i + 1
	}

	valid := func(r, c int) bool {
		return 0 <= r && r < row && 0 <= c && c < col
	}
	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	seen := make([][]bool, row)
	for i := range len(seen) {
		seen[i] = make([]bool, col)
	}
	check := func(day int) bool {
		var q [][]int
		for i := range col {
			if day < water[0][i] {
				q = append(q, []int{0, i})
			}
		}
		// clear seen
		for i := range row {
			for j := range col {
				seen[i][j] = false
			}
		}
		for len(q) > 0 {
			cur := q[0]
			q = q[1:]
			if cur[0] == row-1 {
				return true
			}
			for _, dir := range directions {
				nextRow, nextCol := cur[0]+dir[0], cur[1]+dir[1]
				if valid(nextRow, nextCol) && !seen[nextRow][nextCol] && water[nextRow][nextCol] > day {
					seen[nextRow][nextCol] = true
					q = append(q, []int{nextRow, nextCol})
				}
			}
		}
		return false
	}

	left := 0
	right := len(cells) + 1
	for left < right {
		mid := left + (right-left)/2
		if check(mid) {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left - 1
}
