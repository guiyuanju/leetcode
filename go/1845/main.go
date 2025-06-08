package main

import "fmt"
import "container/heap"

func main() {
}

type SeatManager struct {
	heap *Heap
	n    int
}

func Constructor(n int) SeatManager {
	hv := make([]int, n)
	for i := range n {
		hv[i] = i
	}
	h := Heap(hv)
	heap.Init(&h)

	return SeatManager{
		heap: &h,
		n:    n,
	}
}

func (this *SeatManager) Reserve() int {
	return heap.Pop(this.heap).(int)
}

func (this *SeatManager) Unreserve(seatNumber int) {
	heap.Push(this.heap, seatNumber)
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

/**
 * Your SeatManager object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Reserve();
 * obj.Unreserve(seatNumber);
 */
