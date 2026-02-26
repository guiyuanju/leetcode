package main

import (
	"fmt"
	"reflect"
)

func main() {
	assertEq(true, canFinish(2, [][]int{{1, 0}}))
	assertEq(false, canFinish(2, [][]int{{1, 0}, {0, 1}}))
}

func assertEq(a, b any) {
	if reflect.DeepEqual(a, b) {
		fmt.Printf("Ok: %v\n", a)
	} else {
		fmt.Printf("Failed: %v != %v\n", a, b)
	}
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	ind := make([]int, numCourses)
	g := map[int][]int{}
	for _, p := range prerequisites {
		g[p[1]] = append(g[p[1]], p[0])
		ind[p[0]]++
	}

	q := []int{}
	for i, v := range ind {
		if v == 0 {
			q = append(q, i)
		}
	}

	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		for _, n := range g[cur] {
			ind[n]--
			if ind[n] == 0 {
				q = append(q, n)
			}
		}
	}

	for _, v := range ind {
		if v > 0 {
			return false
		}
	}

	return true
}
