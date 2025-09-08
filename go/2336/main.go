package main

import (
	"container/heap"
	"fmt"
)

func main() {
	s := Constructor()
	s.AddBack(2)
	fmt.Println(s.PopSmallest())
	fmt.Println(s.PopSmallest())
	fmt.Println(s.PopSmallest())
	s.AddBack(1)
	fmt.Println(s.PopSmallest())
	fmt.Println(s.PopSmallest())
	fmt.Println(s.PopSmallest())
}

type SmallestInfiniteSet struct {
	h     Heap
	top   int
	exist map[int]bool
}

func Constructor() SmallestInfiniteSet {
	return SmallestInfiniteSet{nil, 1, map[int]bool{}}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	if this.h.Len() > 0 {
		res := heap.Pop(&this.h).(int)
		this.exist[res] = false
		return res
	}
	this.top++
	return this.top - 1
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	if num < this.top && !this.exist[num] {
		heap.Push(&this.h, num)
		this.exist[num] = true
	}
}

/**
 * Your SmallestInfiniteSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.PopSmallest();
 * obj.AddBack(num);
 */

type Heap []int

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i] < h[j] }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *Heap) Pop() any {
	hd := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return hd
}
