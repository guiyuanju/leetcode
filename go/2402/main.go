package main

import "fmt"

import "container/heap"
import "slices"

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
	counts = make([]int, n)
	ends = make([]int, n)
	slices.SortFunc(meetings, func(a, b []int) int { return a[0] - b[0] })

	empties := new(EmptyHeap)
	usings := new(UsingHeap)
	for i := range n {
		heap.Push(empties, i)
	}

	for _, m := range meetings {
		// using -> empty
		for usings.Len() > 0 {
			r := heap.Pop(usings).(int)
			if ends[r] <= m[0] {
				heap.Push(empties, r)
			} else {
				heap.Push(usings, r)
				break
			}
		}

		if empties.Len() > 0 {
			r := heap.Pop(empties).(int)
			ends[r] = m[1]
			counts[r]++
			heap.Push(usings, r)
		} else {
			r := heap.Pop(usings).(int)
			ends[r] += m[1] - m[0]
			counts[r]++
			heap.Push(usings, r)
		}
	}

	return slices.Index(counts, slices.Max(counts))
}

var counts []int
var ends []int

type UsingHeap []int

func (h UsingHeap) Len() int { return len(h) }
func (h UsingHeap) Less(i, j int) bool {
	return ends[h[i]] < ends[h[j]] || (ends[h[i]] == ends[h[j]] && h[i] < h[j])
}
func (h UsingHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *UsingHeap) Push(x any)   { *h = append(*h, x.(int)) }
func (h *UsingHeap) Pop() any {
	old := *h
	n := len(old)
	res := old[n-1]
	*h = old[:n-1]
	return res
}

type EmptyHeap []int

func (h EmptyHeap) Len() int           { return len(h) }
func (h EmptyHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h EmptyHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *EmptyHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *EmptyHeap) Pop() any {
	old := *h
	n := len(old)
	res := old[n-1]
	*h = old[:n-1]
	return res
}
