package main

import "fmt"
import "container/heap"

func main() {
	nums1 := []int{1, 7, 11}
	nums2 := []int{2, 4, 6}
	k := 3
	fmt.Println(kSmallestPairs(nums1, nums2, k))

	nums1 = []int{1, 1, 2}
	nums2 = []int{1, 2, 3}
	k = 2
	fmt.Println(kSmallestPairs(nums1, nums2, k))
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	h := new(Heap)
	for _, m := range nums1 {
		for _, n := range nums2 {
			heap.Push(h, []int{m, n})
			if h.Len() > k {
				top := heap.Pop(h).([]int)
				if m+n >= top[0]+top[1] {
					break
				}
			}
		}
	}

	var res [][]int
	for h.Len() > 0 {
		res = append(res, heap.Pop(h).([]int))
	}
	return res
}

type Heap [][]int

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i][0]+h[i][1] > h[j][0]+h[j][1] }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x any)        { *h = append(*h, x.([]int)) }
func (h *Heap) Pop() any {
	old := *h
	n := len(old)
	res := old[n-1]
	*h = old[:n-1]
	return res
}
