package main

import (
	"fmt"
	"reflect"
	"slices"
	"strconv"
	"strings"
)

func main() {
	flowers := parse2DArray("[[1,6],[3,7],[9,12],[4,13]]")
	people := []int{2, 3, 7, 11}
	assertEq([]int{1, 2, 2, 2}, (fullBloomFlowers(flowers, people)))

	flowers = parse2DArray("[[1,10],[3,3]]")
	people = []int{3, 3, 2}
	assertEq([]int{2, 2, 1}, (fullBloomFlowers(flowers, people)))

	flowers = parse2DArray("[[32,36],[20,26],[32,32],[43,46],[40,50],[9,10],[19,19],[2,23],[36,37],[38,48],[13,25],[12,48],[21,33],[4,43],[43,49],[35,46],[41,44],[36,44],[40,50],[42,47],[27,50],[7,43],[5,41],[32,35],[24,31],[33,42],[44,47],[32,46],[39,46],[48,50],[10,49],[14,19],[13,20],[41,43],[39,48],[33,44],[23,37],[34,48],[36,36],[6,12],[14,17],[31,34],[28,40],[11,31],[17,50],[31,47],[17,21],[33,49],[20,29],[27,43],[18,47],[46,47],[29,49],[50,50],[5,24],[19,27],[16,24],[18,42],[5,17],[17,26]]")
	people = []int{19, 17, 42, 36, 43, 42, 25, 35, 31, 21, 49, 14, 1, 4, 24, 12, 38, 48, 33, 36, 37, 8, 45, 50, 27, 20, 45, 42, 12, 5, 32, 41, 16, 23, 30, 29, 1, 37, 16, 42, 43, 5, 50, 6, 49, 22, 34, 24, 6}
	assertEq([]int{19, 17, 26, 24, 26, 26, 17, 22, 18, 19, 9, 13, 0, 2, 19, 10, 21, 13, 22, 24, 22, 7, 20, 6, 16, 19, 20, 26, 10, 5, 20, 26, 14, 19, 16, 17, 0, 22, 14, 26, 26, 5, 6, 6, 9, 18, 22, 19, 6}, (fullBloomFlowers(flowers, people)))
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

func fullBloomFlowers(flowers [][]int, people []int) []int {
	starts := make([]int, len(flowers))
	for i, f := range flowers {
		starts[i] = f[0]
	}
	slices.Sort(starts)

	ends := make([]int, len(flowers))
	for i, f := range flowers {
		ends[i] = f[1]
	}
	slices.Sort(ends)

	bsWith := func(x int, xs []int, cmp func(a, b int) bool) int {
		left := 0
		right := len(xs)
		for left < right {
			mid := left + (right-left)/2
			if cmp(xs[mid], x) {
				right = mid
			} else {
				left = mid + 1
			}
		}
		return left
	}

	res := make([]int, len(people))
	for i, p := range people {
		start := bsWith(p, starts, func(a, b int) bool { return a > b }) // right insertion position
		end := bsWith(p, ends, func(a, b int) bool { return a >= b })    // left insertion position
		res[i] = start - end
	}
	return res
}
