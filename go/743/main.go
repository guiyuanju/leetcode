package main

import (
	"container/heap"
	"fmt"
	"math"
)

func main() {
	assertEq(2, networkDelayTime([][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}, 4, 2))
	assertEq(1, networkDelayTime([][]int{{1, 2, 1}}, 2, 1))
	assertEq(-1, networkDelayTime([][]int{{1, 2, 1}}, 2, 2))
}

func assertEq[T comparable](a, b T) {
	if a == b {
		fmt.Printf("✓ %v == %v\n", a, b)
	} else {
		fmt.Printf("✗ %v != %v\n", a, b)
	}
}

func networkDelayTime(times [][]int, n int, k int) int {
	g := map[int][][]int{}
	for _, t := range times {
		g[t[0]] = append(g[t[0]], t[1:])
	}

	dists := make([]int, n+1)
	for i := range dists {
		dists[i] = math.MaxInt
	}
	dists[k] = 0

	h := Heap{[]int{k}, dists}
	for h.Len() > 0 {
		cur := heap.Pop(&h).(int)
		for _, nei := range g[cur] {
			if dists[cur]+nei[1] < dists[nei[0]] {
				dists[nei[0]] = dists[cur] + nei[1]
				heap.Push(&h, nei[0])
			}
		}
	}

	var res int
	for i := 1; i < len(dists); i++ {
		dist := dists[i]
		if dist == math.MaxInt {
			return -1
		}
		res = max(res, dist)
	}

	return res
}

type Heap struct {
	nodes []int
	dists []int
}

func (h Heap) Len() int           { return len(h.nodes) }
func (h Heap) Less(i, j int) bool { return h.dists[h.nodes[i]] < h.dists[h.nodes[j]] }
func (h *Heap) Swap(i, j int)     { (*h).nodes[i], (*h).nodes[j] = (*h).nodes[j], (*h).nodes[i] }
func (h *Heap) Push(x any)        { (*h).nodes = append((*h).nodes, x.(int)) }
func (h *Heap) Pop() any {
	oldLen := len((*h).nodes)
	res := (*h).nodes[oldLen-1]
	(*h).nodes = (*h).nodes[:oldLen-1]
	return res
}
