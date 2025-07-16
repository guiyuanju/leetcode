package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func main() {
	grid := parse2DArray("[[0,2],[1,3]]")
	assertEq(3, swimInWater(grid))

	grid = parse2DArray("[[0,1,2,3,4],[24,23,22,21,5],[12,13,14,15,16],[11,17,18,19,20],[10,9,8,7,6]]")
	assertEq(16, swimInWater(grid))
}

func assertEq(a, b any) {
	if reflect.DeepEqual(a, b) {
		fmt.Printf("Ok: %v\n", a)
	} else {
		fmt.Printf("Failed: %v != %v\n", a, b)
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

func swimInWater(grid [][]int) int {
	valid := func(r, c int) bool {
		return 0 <= r && r < len(grid) && 0 <= c && c < len(grid[0])
	}
	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	check := func(t int) bool {
		q := [][]int{{0, 0}}
		seen := make([][]bool, len(grid))
		for i := range len(seen) {
			seen[i] = make([]bool, len(grid[0]))
		}
		for len(q) > 0 {
			cur := q[0]
			q = q[1:]
			if cur[0] == len(grid)-1 && cur[1] == len(grid[0])-1 {
				return true
			}
			for _, dir := range directions {
				nextRow, nextCol := cur[0]+dir[0], cur[1]+dir[1]
				if valid(nextRow, nextCol) && !seen[nextRow][nextCol] && grid[cur[0]][cur[1]] <= t && grid[nextRow][nextCol] <= t {
					seen[nextRow][nextCol] = true
					q = append(q, []int{nextRow, nextCol})
				}
			}
		}
		return false
	}

	var left, right int
	for _, r := range grid {
		for _, h := range r {
			right = max(right, h)
		}
	}
	right++

	for left < right {
		mid := left + (right-left)/2
		if check(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
