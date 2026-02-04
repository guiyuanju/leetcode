package main

import (
	"fmt"
	"reflect"
	"slices"
	"strconv"
	"strings"
)

func main() {
	rect := parse2DArray("[[1,2],[2,3],[2,5]]")
	points := parse2DArray("[[2,1],[1,4]]")
	assertEq([]int{2, 1}, countRectangles(rect, points))

	rect = parse2DArray("[[1,1],[2,2],[3,3]]")
	points = parse2DArray("[[1,3],[1,1]]")
	assertEq([]int{1, 3}, countRectangles(rect, points))

	rect = parse2DArray("[[4,7],[4,9],[8,5],[6,2],[6,4]]")
	points = parse2DArray("[[4,2],[5,6]]")
	assertEq([]int{5, 0}, countRectangles(rect, points))
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

func countRectangles(rectangles [][]int, points [][]int) []int {
	lengths := make([][]int, 101)
	for _, r := range rectangles {
		lengths[r[1]] = append(lengths[r[1]], r[0])
	}

	for _, l := range lengths {
		slices.Sort(l)
	}

	query := func(w int, xs []int) int {
		i := 0
		j := len(xs)
		for i < j {
			mid := i + (j-i)/2
			if xs[mid] < w {
				i = mid + 1
			} else {
				j = mid
			}
		}
		return len(xs) - i
	}

	res := make([]int, len(points))
	for i, p := range points {
		for j := p[1]; j <= 100; j++ {
			res[i] += query(p[0], lengths[j])
		}
	}

	return res
}
