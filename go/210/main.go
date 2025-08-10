package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(findOrder(2, [][]int{{1, 0}}))
	fmt.Println(findOrder(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}))
	fmt.Println(findOrder(1, [][]int{}))
}

func findOrder(numCourses int, prerequisites [][]int) []int {
	g := make(map[int][]int, numCourses)
	degree := make([]int, numCourses)
	for _, p := range prerequisites {
		g[p[0]] = append(g[p[0]], p[1])
		degree[p[1]]++
	}

	q := make([]int, 0, numCourses)
	for n, d := range degree {
		if d == 0 {
			q = append(q, n)
		}
	}

	res := make([]int, 0, numCourses)
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		res = append(res, cur)
		for _, nei := range g[cur] {
			degree[nei]--
			if degree[nei] == 0 {
				q = append(q, nei)
			}
		}
	}

	if len(res) < numCourses {
		return nil
	}
	slices.Reverse(res)
	return res
}
