package main

import (
	"container/heap"
	"fmt"
	"slices"
)

func main() {
	n := 2
	meetings := [][]int{{0, 10}, {1, 5}, {2, 7}, {3, 4}}
	assertEq(0, mostBooked(n, meetings))

	n = 3
	meetings = [][]int{{1, 20}, {2, 10}, {3, 5}, {4, 9}, {6, 8}}
	assertEq(1, mostBooked(n, meetings))

	n = 100
	meetings = [][]int{{0, 1}}
	assertEq(0, mostBooked(n, meetings))

	n = 4
	meetings = [][]int{{18, 19}, {3, 12}, {17, 19}, {2, 13}, {7, 10}}
	assertEq(0, mostBooked(n, meetings))
	// [2, 13, 0] [3, 12, 1] [7, 10, 2] [17, 19, 3]
	// 1 1 1 1

	n = 3
	meetings = [][]int{{44, 50}, {22, 37}, {46, 49}, {35, 45}, {11, 21}, {31, 32}, {16, 45}}
	assertEq(0, mostBooked(n, meetings))
}

func assertEq[T comparable](a, b T) {
	if a != b {
		fmt.Printf("Failed: %v != %v\n", a, b)
	} else {
		fmt.Printf("Ok: %v\n", a)
	}
}

func mostBooked(n int, meetings [][]int) int {
	count := make([]int, n)
	slices.SortFunc(meetings, func(a, b []int) int {
		return a[0] - b[0]
	})

	unused := Heap([][]int{})
	used := Heap([][]int{})
	for i := range n {
		unused = append(unused, []int{i, i})
	}

	for _, m := range meetings {
		for used.Len() > 0 && used[0][0] <= m[0] {
			cur := heap.Pop(&used).([]int)
			heap.Push(&unused, []int{cur[1], cur[1]})
		}
		if unused.Len() > 0 {
			cur := heap.Pop(&unused).([]int)
			count[cur[0]]++
			cur[0] = m[1]
			heap.Push(&used, cur)
		} else {
			used[0][0] += m[1] - m[0]
			count[used[0][1]]++
			heap.Fix(&used, 0)
		}
	}

	var res, hi int
	for i, c := range count {
		if c > hi {
			hi = c
			res = i
		}
	}
	return res
}

type Heap [][]int

func (h Heap) Len() int { return len(h) }
func (h Heap) Less(i, j int) bool {
	if h[i][0] == h[j][0] {
		return h[i][1] < h[j][1]
	}
	return h[i][0] < h[j][0]
}
func (h Heap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x any)   { *h = append(*h, x.([]int)) }
func (h *Heap) Pop() any {
	hd := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return hd
}
