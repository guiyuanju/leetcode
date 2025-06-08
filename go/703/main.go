package main

import "fmt"
import "container/heap"

func main() {
	h := Constructor(3, []int{4, 5, 8, 2})
	fmt.Println(h.Add(3))
	fmt.Println(h.Add(5))
	fmt.Println(h.Add(10))
	fmt.Println(h.Add(9))
	fmt.Println(h.Add(4))
}

type KthLargest struct {
	heap *Heap
	k    int
}

func Constructor(k int, nums []int) KthLargest {
	h := new(Heap)
	res := KthLargest{
		heap: h,
		k:    k,
	}
	for _, n := range nums {
		res.Add(n)
	}
	return res
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this.heap, val)
	if this.heap.Len() > this.k {
		heap.Pop(this.heap)
	}
	return (*this.heap)[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */

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
