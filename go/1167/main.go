package main

import "fmt"
import "container/heap"

func main() {
	sticks := []int{2, 4, 3}
	assertEq(14, connectSticks(sticks))

	sticks = []int{1, 8, 3, 5}
	assertEq(30, connectSticks(sticks))

	sticks = []int{5}
	assertEq(0, connectSticks(sticks))
}

func assertEq[T comparable](a, b T) {
	if a != b {
		fmt.Printf("Failed: %v != %v\n", a, b)
	} else {
		fmt.Printf("Ok: %v\n", a)
	}
}

func connectSticks(sticks []int) int {
	hv := Heap(sticks)
	h := &hv
	heap.Init(h)
	var cost int
	for h.Len() > 1 {
		top1 := heap.Pop(h).(int)
		top2 := heap.Pop(h).(int)
		heap.Push(h, top1+top2)
		cost += top1 + top2
	}
	return cost
}

type Heap []int

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i] < h[j] }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *Heap) Pop() any {
	old := *h
	n := len(old)
	res := old[n-1]
	*h = old[:n-1]
	return res
}
