package main

import (
	"container/heap"
	"fmt"
)

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
	var res [][]int
	seen := map[[2]int]bool{}
	h := Heap([][3]int{{0, 0, nums1[0] + nums2[0]}})
	seen[[2]int{0, 0}] = true

	for len(res) < k && h.Len() > 0 {
		cur := heap.Pop(&h).([3]int)
		i := cur[0]
		j := cur[1]
		res = append(res, []int{nums1[i], nums2[j]})

		if i < len(nums1)-1 && !seen[[2]int{i + 1, j}] {
			seen[[2]int{i + 1, j}] = true
			heap.Push(&h, [3]int{i + 1, j, nums1[i+1] + nums2[j]})
		}

		if j < len(nums2)-1 && !seen[[2]int{i, j + 1}] {
			seen[[2]int{i, j + 1}] = true
			heap.Push(&h, [3]int{i, j + 1, nums1[i] + nums2[j+1]})
		}

		if i < len(nums1)-1 && j < len(nums2)-1 && !seen[[2]int{i + 1, j + 1}] {
			seen[[2]int{i + 1, j + 1}] = true
			heap.Push(&h, [3]int{i + 1, j + 1, nums1[i+1] + nums2[j+1]})
		}
	}

	return res
}

type Heap [][3]int

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i][2] < h[j][2] }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x any)        { *h = append(*h, x.([3]int)) }
func (h *Heap) Pop() any {
	hd := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return hd
}
