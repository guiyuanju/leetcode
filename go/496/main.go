package main

import "fmt"

func main() {
	fmt.Println(nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2}))
	fmt.Println(nextGreaterElement([]int{2, 4}, []int{1, 2, 3, 4}))
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	m := map[int]int{}
	s := []int{}
	for _, n := range nums2 {
		for len(s) > 0 && s[len(s)-1] < n {
			popped := s[len(s)-1]
			s = s[:len(s)-1]
			m[popped] = n
		}
		s = append(s, n)
	}

	var res []int
	for _, n := range nums1 {
		ng := -1
		if v, ok := m[n]; ok {
			ng = v
		}
		res = append(res, ng)
	}
	return res
}
