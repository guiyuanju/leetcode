package main

import "fmt"

func main() {
	fmt.Println(canFinish(2, [][]int{{1, 0}}))
	fmt.Println(canFinish(2, [][]int{{1, 0}, {0, 1}}))
}

type color int

const (
	white = iota
	gray
	black
)

func canFinish3StatesDFS(numCourses int, prerequisites [][]int) bool {
	g := make(map[int][]int, numCourses)
	for _, p := range prerequisites {
		g[p[0]] = append(g[p[0]], p[1])
	}

	state := make([]color, numCourses)
	for i := range state {
		state[i] = white
	}

	var dfs func(n int) bool
	dfs = func(n int) bool {
		if state[n] == gray {
			return false
		}
		if state[n] == black {
			return true
		}

		state[n] = gray
		for _, nei := range g[n] {
			if !dfs(nei) {
				return false
			}
		}
		state[n] = black
		return true
	}

	for n, c := range state {
		if c == white {
			if !dfs(n) {
				return false
			}
		}
	}

	return true
}

// Kahn's algorithm
func canFinish(numCourses int, prerequisites [][]int) bool {
	g := make(map[int][]int, numCourses)
	indegree := make([]int, numCourses)
	for _, p := range prerequisites {
		g[p[0]] = append(g[p[0]], p[1])
		indegree[p[1]]++
	}

	q := make([]int, 0, numCourses)
	for n, d := range indegree {
		if d == 0 {
			q = append(q, n)
		}
	}

	var count int
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		count++
		for _, nei := range g[cur] {
			indegree[nei]--
			if indegree[nei] == 0 {
				q = append(q, nei)
			}
		}
	}

	return count == numCourses
}
