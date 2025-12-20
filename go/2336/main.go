package main

import (
	"container/heap"
	"fmt"
	"reflect"
)

func main() {
	s := Constructor()
	s.AddBack(2)
	assertEq(1, s.PopSmallest())
	assertEq(2, s.PopSmallest())
	assertEq(3, s.PopSmallest())
	s.AddBack(1)
	assertEq(1, s.PopSmallest())
	assertEq(4, s.PopSmallest())
	assertEq(5, s.PopSmallest())
}

func assertEq(a, b any) {
	if reflect.DeepEqual(a, b) {
		fmt.Printf("Ok: %v\n", a)
	} else {
		fmt.Printf("Failed: %v != %v\n", a, b)
	}
}

type SmallestInfiniteSet struct {
	left      Heap
	leftExist map[int]bool
	rightHead int
}

func Constructor() SmallestInfiniteSet {
	return SmallestInfiniteSet{Heap{}, map[int]bool{}, 1}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	if len(this.left) > 0 {
		res := heap.Pop(&this.left).(int)
		this.leftExist[res] = false
		return res
	}
	this.leftExist[this.rightHead] = false
	this.rightHead++
	return this.rightHead - 1
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	if num >= this.rightHead || this.leftExist[num] {
		return
	}
	heap.Push(&this.left, num)
	this.leftExist[num] = true
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
