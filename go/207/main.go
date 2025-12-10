package main

import "fmt"

func main() {
	fmt.Println(canFinish(2, [][]int{{1, 0}}))
	fmt.Println(canFinish(2, [][]int{{1, 0}, {0, 1}}))
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make(map[int][]int, numCourses)
	indegrees := make([]int, numCourses)
	for _, p := range prerequisites {
		graph[p[1]] = append(graph[p[1]], p[0])
		indegrees[p[0]]++
	}

	queue := []int{}
	for i, in := range indegrees {
		if in == 0 {
			queue = append(queue, i)
		}
	}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		for _, nei := range graph[cur] {
			indegrees[nei]--
			if indegrees[nei] == 0 {
				queue = append(queue, nei)
			}
		}
	}

	for _, in := range indegrees {
		if in != 0 {
			return false
		}
	}

	return true
}
