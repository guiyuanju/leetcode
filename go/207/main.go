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
	g := map[int][]int{}
	indgree := make([]int, numCourses)
	for _, p := range prerequisites {
		indgree[p[1]]++
		g[p[0]] = append(g[p[0]], p[1])
	}

	q := []int{}
	for i, ind := range indgree {
		if ind == 0 {
			q = append(q, i)
		}
	}

	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		for _, nei := range g[cur] {
			indgree[nei]--
			if indgree[nei] == 0 {
				q = append(q, nei)
			}
		}
	}

	for _, ind := range indgree {
		if ind > 0 {
			return false
		}
	}

	return true
}
