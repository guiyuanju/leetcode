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

	mf = Constructor()
	mf.AddNum(1)
	mf.AddNum(2)
	assertEq(1.5, mf.FindMedian())
	mf.AddNum(3)
	assertEq(2, mf.FindMedian())
}

func assertEq[T comparable](a, b T) {
	if a != b {
		fmt.Printf("Failed: %v != %v\n", a, b)
	} else {
		fmt.Printf("Ok: %v\n", a)
	}
}

type MedianFinder struct {
	lo, hi Heap
}

func Constructor() MedianFinder {
	return MedianFinder{}
}

func (this *MedianFinder) AddNum(num int) {
	heap.Push(&this.lo, -num)
	if len(this.lo) > len(this.hi) {
		heap.Push(&this.hi, -heap.Pop(&this.lo).(int))
	}
	if len(this.lo) < len(this.hi) {
		heap.Push(&this.lo, -heap.Pop(&this.hi).(int))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if len(this.lo) > len(this.hi) {
		return float64(-this.lo[0])
	}
	return float64(-this.lo[0]+this.hi[0]) / 2
}

type Heap []int

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i] < h[j] }
func (h *Heap) Swap(i, j int)     { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *Heap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *Heap) Pop() any {
	l := len(*h) - 1
	res := (*h)[l]
	*h = (*h)[:l]
	return res
}
