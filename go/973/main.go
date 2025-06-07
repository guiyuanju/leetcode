package main

import "fmt"
import "math"
import "container/heap"

func main() {
	points := [][]int{{1, 3}, {-2, 2}}
	k := 1
	fmt.Println(kClosest(points, k))

	points = [][]int{{3, 3}, {5, -1}, {-2, 4}}
	k = 2
	fmt.Println(kClosest(points, k))
}

func kClosest(points [][]int, k int) [][]int {
	h := new(Heap)
	for _, p := range points {
		item := makeItem(p)
		heap.Push(h, &item)
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	var res [][]int
	for _, p := range *h {
		res = append(res, p.point)
	}
	return res
}

type Heap []*Item
type Item struct {
	point []int
	dist  float64
}

func makeItem(p []int) Item {
	return Item{
		point: p,
		dist:  math.Sqrt(float64(p[0]*p[0] + p[1]*p[1])),
	}
}

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i].dist > h[j].dist }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x any)        { *h = append(*h, x.(*Item)) }
func (h *Heap) Pop() any {
	old := *h
	n := len(old)
	res := old[n-1]
	*h = old[:n-1]
	return res
}
