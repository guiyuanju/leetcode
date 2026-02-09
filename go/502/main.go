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
	items := make([]Item, len(profits))
	for i := range profits {
		items[i] = Item{profits[i], capital[i]}
	}

	slices.SortFunc(items, func(a, b Item) int { return a.c - b.c })

	h := Heap{}
	curCap := w
	curProjectCount := 0
	curCapIdx := 0
	for curProjectCount < k {
		for ; curCapIdx < len(items) && items[curCapIdx].c <= curCap; curCapIdx++ {
			heap.Push(&h, items[curCapIdx])
		}
		if h.Len() == 0 {
			break
		}
		item := heap.Pop(&h).(Item)
		curCap += item.p
		curProjectCount++
	}

	return curCap
}

type Item struct {
	p, c int
}

type Heap []Item

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i].p > h[j].p }
func (h *Heap) Swap(i, j int)     { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *Heap) Push(x any)        { *h = append(*h, x.(Item)) }
func (h *Heap) Pop() any {
	oldLen := len(*h)
	res := (*h)[oldLen-1]
	*h = (*h)[:oldLen-1]
	return res
}
