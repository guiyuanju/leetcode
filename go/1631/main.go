package main

import (
	"container/heap"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	heights := parse2DArray("[[1,2,2],[3,8,2],[5,3,5]]")
	assertEq(2, minimumEffortPath(heights))

	heights = parse2DArray("[[1,2,3],[3,8,4],[5,3,5]]")
	assertEq(1, minimumEffortPath(heights))

	heights = parse2DArray("[[1,2,1,1,1],[1,2,1,2,1],[1,2,1,2,1],[1,2,1,2,1],[1,1,1,2,1]]")
	assertEq(0, minimumEffortPath(heights))
}

func assertEq[T comparable](a, b T) {
	if a != b {
		fmt.Printf("Failed: %v != %v\n", a, b)
	} else {
		fmt.Printf("Ok: %v\n", a)
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

// bianry search
func minimumEffortPath(heights [][]int) int {
	valid := func(r, c int) bool {
		return 0 <= r && r < len(heights) && 0 <= c && c < len(heights[0])
	}
	check := func(effort int) bool {
		seen := map[[2]int]bool{}
		seen[[2]int{0, 0}] = true
		var dfs func(r, c int) bool
		dfs = func(r, c int) bool {
			if r == len(heights)-1 && c == len(heights[0])-1 {
				return true
			}
			neighbors := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
			for _, nei := range neighbors {
				nextRow, nextCol := r+nei[0], c+nei[1]
				if valid(nextRow, nextCol) && !seen[[2]int{nextRow, nextCol}] && abs(heights[nextRow][nextCol]-heights[r][c]) <= effort {
					seen[[2]int{nextRow, nextCol}] = true
					if dfs(nextRow, nextCol) {
						return true
					}
				}
			}
			return false
		}
		return dfs(0, 0)
	}

	var left, right int
	for i := range len(heights) {
		for j := range len(heights[i]) {
			right = max(right, heights[i][j])
		}
	}

	for left < right {
		mid := left + (right-left)/2
		if check(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Dijkstra
func minimumEffortPath2(heights [][]int) int {
	neighbors := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	valid := func(r, c int) bool {
		return 0 <= r && r < len(heights) && 0 <= c && c < len(heights[0])
	}
	height := func(r0, c0, r1, c1 int) int {
		return abs(heights[r0][c0] - heights[r1][c1])
	}
	dist := make([][]int, len(heights))
	for i := range dist {
		dist[i] = make([]int, len(heights[0]))
		for j := range dist[i] {
			dist[i][j] = math.MaxInt
		}
	}
	dist[0][0] = 0

	h := Heap([]Item{{0, 0, 0}})
	for h.Len() > 0 {
		cur := heap.Pop(&h).(Item)
		if cur.dist > dist[cur.r][cur.c] {
			continue
		}
		for _, nei := range neighbors {
			nextRow, nextCol := cur.r+nei[0], cur.c+nei[1]
			if !valid(nextRow, nextCol) {
				continue
			}
			neiDist := max(cur.dist, height(nextRow, nextCol, cur.r, cur.c))
			if neiDist < dist[nextRow][nextCol] {
				dist[nextRow][nextCol] = neiDist
				heap.Push(&h, Item{nextRow, nextCol, neiDist})
			}
		}
	}

	return dist[len(heights)-1][len(heights[0])-1]
}

type Item struct {
	r, c int
	dist int
}

type Heap []Item

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i].dist < h[j].dist }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x any)        { *h = append(*h, x.(Item)) }
func (h *Heap) Pop() any {
	old := *h
	length := len(old)
	res := old[length-1]
	*h = old[:length-1]
	return res
}
