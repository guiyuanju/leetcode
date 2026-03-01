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
	slices.SortFunc(meetings, func(a, b []int) int { return a[0] - b[0] })
	empty := Heap([]Ele{})
	for i := range n {
		empty = append(empty, Ele{0, i})
	}
	heap.Init(&empty)
	using := Heap([]Ele{})
	count := make([]int, n)
	for _, m := range meetings {
		for len(using) > 0 && using[0].end <= m[0] {
			room := heap.Pop(&using).(Ele)
			room.end = 0
			heap.Push(&empty, room)
		}
		var target Ele
		if len(empty) == 0 {
			target = heap.Pop(&using).(Ele)
			target.end += m[1] - m[0]
		} else {
			target = heap.Pop(&empty).(Ele)
			target.end = m[1]
		}
		count[target.num]++
		heap.Push(&using, target)
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

type Ele struct {
	end int
	num int
}

type Heap []Ele

func (h Heap) Len() int { return len(h) }
func (h Heap) Less(i, j int) bool {
	if h[i].end == h[j].end {
		return h[i].num < h[j].num
	}
	return h[i].end < h[j].end
}
func (h Heap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x any)   { *h = append(*h, x.(Ele)) }
func (h *Heap) Pop() any {
	hd := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return hd
}
