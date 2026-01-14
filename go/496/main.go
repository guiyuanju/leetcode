package main

import (
	"fmt"
	"reflect"
)

func main() {
	assertEq([]int{-1, 3, -1}, nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2}))
	assertEq([]int{3, -1}, nextGreaterElement([]int{2, 4}, []int{1, 2, 3, 4}))
}

func assertEq(a, b any) {
	if reflect.DeepEqual(a, b) {
		fmt.Printf("Ok: %v\n", a)
	} else {
		fmt.Printf("Failed: %v != %v\n", a, b)
	}
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

	res := make([]int, len(nums1))
	for i := range res {
		if v, ok := m[nums1[i]]; ok {
			res[i] = v
		} else {
			res[i] = -1
		}
	}

	return res
}
