package main

import (
	"container/heap"
	"fmt"
	"reflect"
)

func main() {
	nums1 := []int{1, 7, 11}
	nums2 := []int{2, 4, 6}
	k := 3
	assertEq([][]int{{1, 2}, {1, 4}, {1, 6}}, kSmallestPairs(nums1, nums2, k))

	nums1 = []int{1, 1, 2}
	nums2 = []int{1, 2, 3}
	k = 2
	assertEq([][]int{{1, 1}, {1, 1}}, kSmallestPairs(nums1, nums2, k))
}

func assertEq(a, b any) {
	if reflect.DeepEqual(a, b) {
		fmt.Printf("Ok: %v\n", a)
	} else {
		fmt.Printf("Failed: %v != %v\n", a, b)
	}
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	var res [][]int
	seen := map[[2]int]bool{}
	seen[[2]int{0, 0}] = true

	h := Heap{[3]int{0, 0, nums1[0] + nums2[0]}}
	for k > 0 {
		cur := heap.Pop(&h).([3]int)
		if cur[0]+1 < len(nums1) && !seen[[2]int{cur[0] + 1, cur[1]}] {
			heap.Push(&h, [3]int{cur[0] + 1, cur[1], nums1[cur[0]+1] + nums2[cur[1]]})
			seen[[2]int{cur[0] + 1, cur[1]}] = true
		}
		if cur[1]+1 < len(nums2) && !seen[[2]int{cur[0], cur[1] + 1}] {
			heap.Push(&h, [3]int{cur[0], cur[1] + 1, nums1[cur[0]] + nums2[cur[1]+1]})
			seen[[2]int{cur[0], cur[1] + 1}] = true
		}
		res = append(res, []int{nums1[cur[0]], nums2[cur[1]]})
		k--
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
