package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(maximumProduct([]int{0, 4}, 5))
	fmt.Println(maximumProduct([]int{6, 3, 3, 2}, 2))
}

func maximumProduct(nums []int, k int) int {
	h := Heap(nums)
	heap.Init(&h)
	for k > 0 {
		v := heap.Pop(&h).(int)
		heap.Push(&h, v+1)
		k--
	}

	res := 1
	for h.Len() > 0 {
		res *= heap.Pop(&h).(int)
		res %= 1e9 + 7
	}

	return res
}

type Heap []int

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i] < h[j] }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *Heap) Pop() any {
	old := *h
	length := len(old)
	res := old[length-1]
	*h = old[:length-1]
	return res
}
