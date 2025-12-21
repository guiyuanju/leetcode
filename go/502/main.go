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
	sortedCapital := make([]int, len(capital))
	for i := range len(sortedCapital) {
		sortedCapital[i] = i
	}
	slices.SortFunc(sortedCapital, func(i, j int) int { return capital[i] - capital[j] })

	var i int
	h := Heap{}
	curCap := w
	for range k {
		for i < len(sortedCapital) && capital[sortedCapital[i]] <= curCap {
			heap.Push(&h, profits[sortedCapital[i]])
			i++
		}
		if len(h) > 0 {
			curCap += heap.Pop(&h).(int)
		}
	}

	return curCap
}

type Heap []int

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i] > h[j] }
func (h *Heap) Swap(i, j int)     { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *Heap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *Heap) Pop() any {
	l := len(*h) - 1
	res := (*h)[l]
	*h = (*h)[:l]
	return res
}
