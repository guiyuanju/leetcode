package main

import (
	"container/heap"
	"fmt"
	"slices"
)

func main() {
	k := 2
	w := 0
	profits := []int{1, 2, 3}
	capital := []int{0, 1, 1}
	assertEq(4, findMaximizedCapital(k, w, profits, capital))

	assertEq(6, findMaximizedCapital(3, 0, []int{1, 2, 3}, []int{0, 1, 2}))

	k = 10
	w = 0
	profits = []int{1, 2, 3}
	capital = []int{0, 1, 2}
	assertEq(6, findMaximizedCapital(k, w, profits, capital))
}

func assertEq[T comparable](a, b T) {
	if a != b {
		fmt.Printf("Failed: %v != %v\n", a, b)
	} else {
		fmt.Printf("Ok: %v\n", a)
	}
}

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	var maxPros Heap
	maxPros = Heap{nil, func(i, j int) bool { return profits[maxPros.vals[i]] > profits[maxPros.vals[j]] }}

	indices := make([]int, len(profits))
	for i := range indices {
		indices[i] = i
	}
	slices.SortFunc(indices, func(i, j int) int { return capital[i] - capital[j] })

	res := w
	var i int
	for range k {
		for i < len(indices) && capital[indices[i]] <= res {
			heap.Push(&maxPros, indices[i])
			i++
		}
		if maxPros.Len() == 0 {
			return res
		}
		res += profits[heap.Pop(&maxPros).(int)]
	}
	return res
}

type Heap struct {
	vals []int
	less func(i, j int) bool
}

func (h Heap) Len() int           { return len(h.vals) }
func (h Heap) Less(i, j int) bool { return h.less(i, j) }
func (h Heap) Swap(i, j int)      { h.vals[i], h.vals[j] = h.vals[j], h.vals[i] }
func (h *Heap) Push(x any)        { (*h).vals = append((*h).vals, x.(int)) }
func (h *Heap) Pop() any {
	hd := (*h).vals[len((*h).vals)-1]
	(*h).vals = (*h).vals[:len((*h).vals)-1]
	return hd
}
