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
		ind[p[1]]++
		g[p[0]] = append(g[p[0]], p[1])
	}

	q := []int{}
	for i, d := range ind {
		if d == 0 {
			q = append(q, i)
		}
	}

	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		for _, nei := range g[cur] {
			ind[nei]--
			if ind[nei] == 0 {
				q = append(q, nei)
			}
		}
	}

	for _, d := range ind {
		if d > 0 {
			return false
		}
	}

	return true
}
