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

	mf = Constructor()
	mf.AddNum(-1)
	assertEq(-1, mf.FindMedian())
	mf.AddNum(-2)
	assertEq(-1.5, mf.FindMedian())
	mf.AddNum(-3)
	assertEq(-2.0, mf.FindMedian())
	mf.AddNum(-4)
	assertEq(-2.5, mf.FindMedian())
	mf.AddNum(-5)
	assertEq(-3.0, mf.FindMedian())
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
	hd := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return hd
}

type MedianFinder struct {
	left, right Heap
}

func Constructor() MedianFinder {
	return MedianFinder{}
}

func (this *MedianFinder) AddNum(num int) {
	heap.Push(&(this.right), num)
	heap.Push(&(this.left), -heap.Pop(&(this.right)).(int))
	if this.left.Len()-this.right.Len() == 2 {
		heap.Push(&(this.right), -heap.Pop(&(this.left)).(int))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.left.Len() > this.right.Len() {
		return -float64(this.left[0])
	}
	return (-float64(this.left[0]) + float64(this.right[0])) / 2
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
