package main

import "fmt"
import "container/heap"

func main() {
	nums := []int{3, 2, 1, 5, 6, 4}
	k := 2
	assertEq(5, findKthLargest(nums, k))

	nums = []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k = 4
	assertEq(4, findKthLargest(nums, k))
}

func assertEq[T comparable](a, b T) {
	if a != b {
		fmt.Printf("Failed: %v != %v\n", a, b)
	} else {
		fmt.Printf("Ok: %v\n", a)
	}
}

func findKthLargest(nums []int, k int) int {
	hv := Heap(nums)
	h := &hv
	heap.Init(h)
	var res int
	for k > 0 {
		res = heap.Pop(h).(int)
		k--
	}
	return res
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
