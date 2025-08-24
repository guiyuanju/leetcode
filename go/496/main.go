package main

import "fmt"

func main() {
	fmt.Println(nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2}))
	fmt.Println(nextGreaterElement([]int{2, 4}, []int{1, 2, 3, 4}))
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	m := map[int]int{}
	mono := []int{}
	for _, n := range nums2 {
		for len(mono) > 0 && n > mono[len(mono)-1] {
			m[mono[len(mono)-1]] = n
			mono = mono[:len(mono)-1]
		}
		mono = append(mono, n)
	}

	res := make([]int, 0, len(nums1))
	for _, n := range nums1 {
		if v, ok := m[n]; ok {
			res = append(res, v)
		} else {
			res = append(res, -1)
		}
	}
	return res
}
