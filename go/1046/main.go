package main

import (
	"container/heap"
	"fmt"
)

func main() {
	stones := []int{2, 7, 4, 1, 8, 1}
	assertEq(1, lastStoneWeight(stones))

	stones = []int{1}
	assertEq(1, lastStoneWeight(stones))
}

func assertEq[T comparable](a, b T) {
	if a != b {
		fmt.Printf("Failed: %v != %v\n", a, b)
	} else {
		fmt.Printf("Ok: %v\n", a)
	}
}

func lastStoneWeight(stones []int) int {
	h := Heap(stones)
	heap.Init(&h)
	for h.Len() > 1 {
		x := heap.Pop(&h).(int)
		y := heap.Pop(&h).(int)
		if x > y {
			heap.Push(&h, x-y)
		}
	}
	if h.Len() > 0 {
		return h[0]
	}
	return 0
}

type Heap []int

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i] > h[j] }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *Heap) Pop() any {
	hd := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return hd
}
