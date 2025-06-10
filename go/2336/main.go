package main

import "fmt"
import "container/heap"

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
	heap *Heap
	cur  int
}

func Constructor() SmallestInfiniteSet {
	return SmallestInfiniteSet{
		heap: new(Heap),
		cur:  1,
	}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	if this.heap.Len() == 0 {
		res := this.cur
		this.cur++
		return res
	}

	res := heap.Pop(this.heap).(int)
	for this.heap.Len() > 0 {
		cur := heap.Pop(this.heap).(int)
		if cur != res {
			heap.Push(this.heap, cur)
			break
		}
	}
	return res
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	if num < this.cur {
		heap.Push(this.heap, num)
	}
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
 * Your SmallestInfiniteSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.PopSmallest();
 * obj.AddBack(num);
 */
