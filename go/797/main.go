package main

import (
	"slices"

	"github.com/guiyuanju/lcutils/assert"
	"github.com/guiyuanju/lcutils/grid"
)

func main() {
	graph := grid.New("[[1,2],[3],[3],[]]")
	assert.Eq(grid.New("[[0,1,3],[0,2,3]]"), allPathsSourceTarget(graph))

	graph = grid.New("[[4,3,1],[3,2,4],[3],[4],[]]")
	assert.Eq(grid.New("[[0,4],[0,3,4],[0,1,3,4],[0,1,2,3,4],[0,1,4]]"), allPathsSourceTarget(graph))
}

func allPathsSourceTarget(graph [][]int) [][]int {
	var res [][]int
	var dfs func([]int, int)
	dfs = func(path []int, i int) {
		if i == len(graph)-1 {
			res = append(res, append([]int(nil), path...))
			return
		}
		for _, nei := range graph[i] {
			if !slices.Contains(path, nei) {
				dfs(append(path, nei), nei)
			}
		}
	}
	dfs([]int{0}, 0)
	return res
}
