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
	graph := make(map[int][][2]int, n)
	for _, t := range times {
		graph[t[0]] = append(graph[t[0]], [2]int{t[1], t[2]})
	}

	dist := make([]int, n+1)

	for i := range dist {
		dist[i] = math.MaxInt
	}

	h := Heap{
		nodes: []int{k},
		less:  func(a, b int) bool { return dist[a] < dist[b] },
	}

	dist[k] = 0

	for h.Len() > 0 {
		cur := heap.Pop(&h).(int)
		for _, nei := range graph[cur] {
			if dist[cur]+nei[1] < dist[nei[0]] {
				dist[nei[0]] = dist[cur] + nei[1]
				heap.Push(&h, nei[0])
			}
		}
	}

	var res int
	for i := 1; i < len(dist); i++ {
		d := dist[i]
		if d == math.MaxInt {
			return -1
		}
		res = max(res, d)
	}
	return res
}

type Heap struct {
	nodes []int
	less  func(int, int) bool
}

func (h Heap) Len() int           { return len(h.nodes) }
func (h Heap) Less(a, b int) bool { return h.less(h.nodes[a], h.nodes[b]) }
func (h *Heap) Swap(a, b int)     { h.nodes[a], h.nodes[b] = h.nodes[b], h.nodes[a] }
func (h *Heap) Push(x any)        { h.nodes = append(h.nodes, x.(int)) }
func (h *Heap) Pop() any {
	oldLen := len(h.nodes)
	res := h.nodes[oldLen-1]
	h.nodes = h.nodes[:oldLen-1]
	return res
}
