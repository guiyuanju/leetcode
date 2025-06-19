package main

import (
	"container/heap"
	"fmt"
	"slices"
)

func main() {
	nums := []int{4, 2, 3}
	k := 1
	assertEq(5, largestSumAfterKNegations2(nums, k))

	nums = []int{3, -1, 0, 2}
	k = 3
	assertEq(6, largestSumAfterKNegations2(nums, k))

	nums = []int{2, -3, -1, 5, -4}
	k = 2
	assertEq(13, largestSumAfterKNegations2(nums, k))

	nums = []int{-2, 9, 9, 8, 4}
	k = 5
	assertEq(32, largestSumAfterKNegations2(nums, k))

	nums = []int{-8, 3, -5, -3, -5, -2}
	k = 6
	assertEq(22, largestSumAfterKNegations2(nums, k))
}

func assertEq[T comparable](a, b T) {
	if a != b {
		fmt.Printf("Failed: %v != %v\n", a, b)
	} else {
		fmt.Printf("Ok: %v\n", a)
	}
}

func largestSumAfterKNegations(nums []int, k int) int {
	slices.Sort(nums)

	for i, n := range nums {
		if k == 0 || n >= 0 {
			break
		}
		if n < 0 {
			nums[i] *= -1
			k--
		}
	}

	if k > 0 {
		slices.Sort(nums)
		if k%2 > 0 {
			nums[0] *= -1
			k = 0
		}
	}

	var res int
	for _, n := range nums {
		res += n
	}
	return res
}

func largestSumAfterKNegations2(nums []int, k int) int {
	h := Heap(nums)
	heap.Init(&h)
	for k > 0 {
		heap.Push(&h, -heap.Pop(&h).(int))
		k--
	}

	var res int
	for _, n := range nums {
		res += n
	}
	return res
}

type Heap []int

func (h Heap) Len() int            { return len(h) }
func (h Heap) Less(i, j int) bool  { return h[i] < h[j] }
func (h Heap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *Heap) Pop() interface{} {
	hd := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return hd
}
