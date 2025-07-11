package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func main() {
	rect := parse2DArray("[[1,2],[2,3],[2,5]]")
	points := parse2DArray("[[2,1],[1,4]]")
	fmt.Println(countRectangles(rect, points))

	rect = parse2DArray("[[1,1],[2,2],[3,3]]")
	points = parse2DArray("[[1,3],[1,1]]")
	fmt.Println(countRectangles(rect, points))

	rect = parse2DArray("[[4,7],[4,9],[8,5],[6,2],[6,4]]")
	points = parse2DArray("[[4,2],[5,6]]")
	fmt.Println(countRectangles(rect, points))
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

func countRectangles(rectangles [][]int, points [][]int) []int {
	widths := make([][]int, 101)
	for _, r := range rectangles {
		widths[r[1]] = append(widths[r[1]], r[0])
	}
	for _, ws := range widths {
		slices.SortFunc(ws, func(a, b int) int { return b - a })
	}

	bs := func(p []int, nums []int) int {
		left := 0
		right := len(nums)
		for left < right {
			mid := left + (right-left)/2
			if nums[mid] < p[0] {
				right = mid
			} else {
				left = mid + 1
			}
		}
		return left
	}

	var res []int
	for _, p := range points {
		var cur int
		for i := p[1]; i <= 100; i++ {
			cur += bs(p, widths[i])
		}
		res = append(res, cur)
	}

	return res
}
