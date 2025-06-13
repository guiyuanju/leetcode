package main

import "fmt"
import "slices"

func main() {
	nums := [][]int{{4, 10, 15, 24, 26}, {0, 9, 12, 20}, {5, 18, 22, 30}}
	fmt.Println(smallestRange(nums))

	nums = [][]int{{1, 2, 3}, {1, 2, 3}, {1, 2, 3}}
	fmt.Println(smallestRange(nums))

	nums = [][]int{
		{-1, 1}, {-2, 2}, {-3, 3}, {-4, 4}, {-5, 5},
		{-6, 6}, {-7, 7}, {-8, 8}, {-9, 9}, {-10, 10},
		{-11, 11}, {-12, 12}, {-13, 13}, {-14, 14}, {-15, 15},
		{-16, 16}, {-17, 17}, {-18, 18}, {-19, 19}, {-20, 20},
		{-21, 21}, {-22, 22}, {-23, 23}, {-24, 24}, {-25, 25},
		{-26, 26}, {-27, 27}, {-28, 28}, {-29, 29}, {-30, 30},
		{-31, 31}, {-32, 32}, {-33, 33}, {-34, 34}, {-35, 35},
		{-36, 36}, {-37, 37}, {-38, 38}, {-39, 39}, {-40, 40},
		{-41, 41}, {-42, 42}, {-43, 43}, {-44, 44}, {-45, 45},
		{-46, 46}, {-47, 47}, {-48, 48}, {-49, 49}, {-50, 50},
		{-51, 51}, {-52, 52}, {-53, 53}, {-54, 54}, {-55, 55},
	}
	fmt.Println(smallestRange(nums))
}

func smallestRange(nums [][]int) []int {
	all := []Item{}
	for i, vs := range nums {
		for _, v := range vs {
			all = append(all, Item{v, i})
		}
	}
	slices.SortFunc(all, func(a, b Item) int {
		return a.value - b.value
	})

	var i, j int
	count := make([]int, len(nums))
	res := []int{-1e5 - 1, 1e5 + 1}
	for ; j < len(all); j++ {
		count[all[j].from]++
		if slices.Index(count, 0) > -1 {
			continue
		}

		for ; i < j && count[all[i].from] > 1; i++ {
			count[all[i].from]--
		}

		if rangeLess(all[i].value, all[j].value, res[0], res[1]) {
			res = []int{all[i].value, all[j].value}
		}
	}

	return res
}

type Item struct {
	value int
	from  int
}

func rangeLess(a, b, c, d int) bool {
	return b-a < d-c || (b-a == d-c && a < c)
}

type Heap []Item

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i].value < h[j].value }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x any)        { *h = append(*h, x.(Item)) }
func (h *Heap) Pop() any {
	old := *h
	n := len(old)
	res := old[n-1]
	*h = old[:n-1]
	return res
}
