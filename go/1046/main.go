package main

import "fmt"
import "container/heap"

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

type Heap []int

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i] > h[j] }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *Heap) Pop() any {
	old := *h
	res := old[len(old)-1]
	*h = old[:len(old)-1]
	return res
}

func lastStoneWeight(stones []int) int {
	h := new(Heap)
	heap.Init(h)
	for _, s := range stones {
		heap.Push(h, s)
	}
	for h.Len() > 1 {
		first := heap.Pop(h).(int)
		second := heap.Pop(h).(int)
		if first > second {
			heap.Push(h, first-second)
		}
	}
	if h.Len() == 0 {
		return 0
	}
	return heap.Pop(h).(int)
}
