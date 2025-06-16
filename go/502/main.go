package main

import "fmt"
import "slices"
import "container/heap"

func main() {
	k := 2
	w := 0
	profits := []int{1, 2, 3}
	capital := []int{0, 1, 1}
	assertEq(4, findMaximizedCapital(k, w, profits, capital))

	k = 3
	w = 0
	profits = []int{1, 2, 3}
	capital = []int{0, 1, 2}
	assertEq(6, findMaximizedCapital(k, w, profits, capital))

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

type Item struct {
	profit, capital int
}

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	var ps []Item
	for i := range len(profits) {
		ps = append(ps, Item{profits[i], capital[i]})
	}

	slices.SortFunc(ps, func(a, b Item) int { return a.capital - b.capital })

	h := new(Heap)
	var i int
	for k > 0 {
		for i < len(ps) && ps[i].capital <= w {
			heap.Push(h, ps[i])
			i++
		}
		if h.Len() == 0 {
			break
		}
		top := heap.Pop(h).(Item)
		w += top.profit
		k--
	}

	return w
}

type Heap []Item

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i].profit > h[j].profit }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x any)        { *h = append(*h, x.(Item)) }
func (h *Heap) Pop() any {
	old := *h
	n := len(old)
	res := old[n-1]
	*h = old[:n-1]
	return res
}
