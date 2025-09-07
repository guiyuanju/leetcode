package main

import "fmt"

func main() {
	fmt.Println(canFinish(2, [][]int{{1, 0}}))
	fmt.Println(canFinish(2, [][]int{{1, 0}, {0, 1}}))
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	g := map[int][]int{}
	for _, p := range prerequisites {
		g[p[0]] = append(g[p[0]], p[1])
	}

	const (
		White = 0
		Gray  = 1
		Black = 2
	)

	states := make([]int, numCourses)
	var dfs func(n int) bool
	dfs = func(n int) bool {
		if states[n] == Gray {
			return false
		}
		if states[n] == Black {
			return true
		}
		states[n] = Gray
		for _, nei := range g[n] {
			if !dfs(nei) {
				return false
			}
		}
		states[n] = Black
		return true
	}

	for n := range numCourses {
		if states[n] == White {
			if !dfs(n) {
				return false
			}
		}
	}
	return true
}
