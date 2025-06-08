package main

import "fmt"
import "container/heap"
import "slices"

func main() {
	words := []string{"i", "love", "leetcode", "i", "love", "coding"}
	k := 2
	fmt.Println(topKFrequent(words, k))

	words = []string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}
	k = 4
	fmt.Println(topKFrequent(words, k))
}

func topKFrequent(words []string, k int) []string {
	clear(count)

	for _, w := range words {
		count[w]++
	}

	h := new(Heap)
	for w, _ := range count {
		heap.Push(h, w)
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	var res []string
	for h.Len() > 0 {
		res = append(res, heap.Pop(h).(string))
	}
	slices.Reverse(res)
	return res
}

var count = map[string]int{}

type Heap []string

func (h Heap) Len() int { return len(h) }
func (h Heap) Less(i, j int) bool {
	return count[h[i]] < count[h[j]] || (count[h[i]] == count[h[j]] && h[i] > h[j])
}
func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *Heap) Push(x any) { *h = append(*h, x.(string)) }
func (h *Heap) Pop() any {
	old := *h
	n := len(old)
	res := old[n-1]
	*h = old[:n-1]
	return res
}
