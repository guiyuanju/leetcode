package main

import "fmt"
import "container/heap"
import "slices"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	k := 4
	x := 3
	fmt.Println(findClosestElements(arr, k, x))

	arr = []int{1, 1, 2, 3, 4, 5}
	k = 4
	x = -1
	fmt.Println(findClosestElements(arr, k, x))
}

func findClosestElements(arr []int, k int, x int) []int {
	h := new(Heap)
	for _, n := range arr {
		item := &Item{
			value: n,
			dist:  abs(n - x),
		}
		heap.Push(h, item)
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	var res []int
	for h.Len() > 0 {
		res = append(res, heap.Pop(h).(*Item).value)
	}
	slices.Sort(res)
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type Heap []*Item
type Item struct {
	value int
	dist  int
}

func (h Heap) Len() int { return len(h) }
func (h Heap) Less(i, j int) bool {
	var res bool
	if h[i].dist < h[j].dist {
		res = true
	} else if h[i].dist == h[j].dist {
		res = h[i].value < h[j].value
	} else {
		res = false
	}
	return !res
}
func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *Heap) Push(x any) {
	*h = append(*h, x.(*Item))
}
func (h *Heap) Pop() any {
	old := *h
	n := len(old)
	res := old[n-1]
	*h = old[:n-1]
	return res
}
