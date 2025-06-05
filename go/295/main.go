package main

import (
	"container/heap"
	"fmt"
)

func main() {
	mf := Constructor()
	mf.AddNum(1)
	mf.AddNum(2)
	assertEq(1.5, mf.FindMedian())
	mf.AddNum(3)
	assertEq(2.0, mf.FindMedian())
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
func (h *Heap) Peek() any {
	res := heap.Pop(h)
	heap.Push(h, res)
	return res
}

type MedianFinder struct {
	left  *Heap
	right *Heap
}

func Constructor() MedianFinder {
	return MedianFinder{new(Heap), new(Heap)}
}

func (this *MedianFinder) AddNum(num int) {
	heap.Push(this.left, -num)
	heap.Push(this.right, -heap.Pop(this.left).(int))
	if this.right.Len() > this.left.Len() {
		heap.Push(this.left, -heap.Pop(this.right).(int))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.left.Len() == this.right.Len() {
		l := -this.left.Peek().(int)
		r := this.right.Peek().(int)
		res := float64(l+r) / 2
		return res
	}
	l := -this.left.Peek().(int)
	return float64(l)
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
