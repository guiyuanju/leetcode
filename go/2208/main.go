package main

import "fmt"
import "container/heap"

func main() {
	nums := []int{5, 19, 8, 1}
	assertEq(3, halveArray(nums))

	nums = []int{3, 8, 20}
	assertEq(3, halveArray(nums))
}

func assertEq[T comparable](a, b T) {
	if a != b {
		fmt.Printf("Failed: %v != %v\n", a, b)
	} else {
		fmt.Printf("Ok: %v\n", a)
	}
}

type Heap []float64

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i] > h[j] }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x any)        { *h = append(*h, x.(float64)) }
func (h *Heap) Pop() any {
	old := *h
	res := old[len(old)-1]
	*h = old[:len(old)-1]
	return res
}

func halveArray(nums []int) int {
	h := new(Heap)
	heap.Init(h)
	var sum float64
	for _, n := range nums {
		sum += float64(n)
		heap.Push(h, float64(n))
	}
	var reduced float64
	var count int
	for reduced < sum/2 {
		largest := heap.Pop(h).(float64)
		reduced += largest / 2
		heap.Push(h, largest/2)
		count++
	}
	return count
}
