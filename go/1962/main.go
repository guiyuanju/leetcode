package main

import (
	"container/heap"
	"fmt"
	"math"
)

func main() {
	piles := []int{5, 4, 9}
	k := 2
	assertEq(12, minStoneSum(piles, k))

	piles = []int{4, 3, 6, 7}
	k = 3
	assertEq(12, minStoneSum(piles, k))
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
	n := len(old)
	res := old[n-1]
	*h = old[:n-1]
	return res
}

func minStoneSum(piles []int, k int) int {
	hv := Heap(piles)
	h := &hv
	heap.Init(h)
	for i := 0; i < k; i++ {
		top := heap.Pop(h).(int)
		new := int(math.Ceil(float64(top) / 2))
		heap.Push(h, new)
	}
	var res int
	for _, v := range piles {
		res += v
	}
	return res
}
