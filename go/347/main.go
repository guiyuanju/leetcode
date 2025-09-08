package main

import (
	"container/heap"
	"fmt"
	"reflect"
)

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	assertEq([]int{1, 2}, topKFrequent(nums, k))

	nums = []int{1}
	k = 1
	assertEq([]int{1}, topKFrequent(nums, k))

	nums = []int{1, 2, 1, 2, 1, 2, 3, 1, 3, 2}
	k = 2
	assertEq([]int{1, 2}, topKFrequent(nums, k))

	nums = []int{3, 0, 1, 0}
	k = 1
	assertEq([]int{0}, topKFrequent(nums, k))
}

func assertEq(a, b any) {
	if reflect.DeepEqual(a, b) {
		fmt.Printf("Ok: %v\n", a)
	} else {
		fmt.Printf("Failed: %v != %v\n", a, b)
	}
}

func topKFrequent(nums []int, k int) []int {
	freq := map[int]int{}
	for _, n := range nums {
		freq[n]++
	}
	h := Heap{nil, func(i, j int) bool {
		return freq[i] < freq[j]
	}}
	for n := range freq {
		heap.Push(&h, n)
		if h.Len() > k {
			heap.Pop(&h)
		}
	}
	return h.vals
}

type Heap struct {
	vals []int
	less func(i, j int) bool
}

func (h Heap) Len() int           { return len(h.vals) }
func (h Heap) Less(i, j int) bool { return h.less(h.vals[i], h.vals[j]) }
func (h Heap) Swap(i, j int)      { h.vals[i], h.vals[j] = h.vals[j], h.vals[i] }
func (h *Heap) Push(x any)        { (*h).vals = append((*h).vals, x.(int)) }
func (h *Heap) Pop() any {
	old := (*h).vals
	n := len(old)
	res := old[n-1]
	(*h).vals = old[:n-1]
	return res
}
