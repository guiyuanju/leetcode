package main

import "fmt"
import "container/heap"

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	fmt.Println(topKFrequent(nums, k))

	nums = []int{1}
	k = 1
	fmt.Println(topKFrequent(nums, k))
}

func topKFrequent(nums []int, k int) []int {
	count := map[int]int{}
	h := new(Heap)
	for _, n := range nums {
		count[n]++
	}

	for i, v := range count {
		heap.Push(h, &Item{i, v})
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	var res []int
	for _, v := range *h {
		item := *v
		res = append(res, item.value)
	}
	return res
}

type Heap []*Item
type Item struct {
	value int
	count int
}

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i].count < h[j].count }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x any)        { *h = append(*h, x.(*Item)) }
func (h *Heap) Pop() any {
	old := *h
	n := len(old)
	res := old[n-1]
	*h = old[:n-1]
	return res
}
