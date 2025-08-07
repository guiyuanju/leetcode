package main

import (
	"container/heap"
	"fmt"
	"math"
	"slices"
)

func main() {
	fmt.Println(networkDelayTime([][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}, 4, 2))
	fmt.Println(networkDelayTime([][]int{{1, 2, 1}}, 2, 1))
	fmt.Println(networkDelayTime([][]int{{1, 2, 1}}, 2, 2))
}

func networkDelayTime(times [][]int, n int, k int) int {
	g := map[int][][]int{}
	for _, e := range times {
		g[e[0]] = append(g[e[0]], []int{e[1], e[2]})
	}

	dist := make([]int, n+1)
	for i := range dist {
		dist[i] = math.MaxInt
	}

	h := Heap([][2]int{{k, 0}})
	for h.Len() > 0 {
		cur := heap.Pop(&h).([2]int)
		if dist[cur[0]] <= cur[1] {
			continue
		}
		dist[cur[0]] = cur[1]
		for _, nei := range g[cur[0]] {
			// important optimization
			if nei[1]+cur[1] < dist[nei[0]] {
				heap.Push(&h, [2]int{nei[0], nei[1] + cur[1]})
			}
		}
	}

	res := slices.Max(dist[1:])
	if res == math.MaxInt {
		return -1
	}
	return res
}

type Heap [][2]int

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i][1] < h[j][1] }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x any)        { *h = append(*h, x.([2]int)) }
func (h *Heap) Pop() any {
	hd := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return hd
}
