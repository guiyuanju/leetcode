package main

import "fmt"
import "container/heap"

func main() {
	costs := []int{17, 12, 10, 2, 7, 2, 11, 20, 8}
	k := 3
	candidates := 4
	assertEq(11, totalCost(costs, k, candidates))

	costs = []int{1, 2, 4, 1}
	k = 3
	candidates = 3
	assertEq(4, totalCost(costs, k, candidates))

	costs = []int{31, 25, 72, 79, 74, 65, 84, 91, 18, 59, 27, 9, 81, 33, 17, 58}
	k = 11
	candidates = 2
	assertEq(423, totalCost(costs, k, candidates))
}

func assertEq[T comparable](a, b T) {
	if a != b {
		fmt.Printf("Failed: %v != %v\n", a, b)
	} else {
		fmt.Printf("Ok: %v\n", a)
	}
}

func totalCost(costs []int, k int, candidates int) int64 {
	var res int64
	h := new(Heap)
	left := 0
	right := len(costs) - 1
	for ; left < candidates; left++ {
		heap.Push(h, &Item{left, costs[left]})
	}
	for ; right >= len(costs)-candidates && right >= left; right-- {
		heap.Push(h, &Item{right, costs[right]})
	}

	for k > 0 {
		cur := heap.Pop(h).(*Item)
		res += int64(cur.cost)
		k--

		if left > right {
			continue
		}
		if cur.index < left {
			heap.Push(h, &Item{left, costs[left]})
			left++
		}
		if cur.index > right {
			heap.Push(h, &Item{right, costs[right]})
			right--
		}
	}

	return res
}

type Heap []*Item
type Item struct {
	index int
	cost  int
}

func (h Heap) Len() int { return len(h) }
func (h Heap) Less(i, j int) bool {
	return h[i].cost < h[j].cost || (h[i].cost == h[j].cost && h[i].index < h[j].index)
}
func (h Heap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x any)   { *h = append(*h, x.(*Item)) }
func (h *Heap) Pop() any {
	old := *h
	n := len(old)
	res := old[n-1]
	*h = old[:n-1]
	return res
}
