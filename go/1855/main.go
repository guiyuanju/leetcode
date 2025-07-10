package main

import "fmt"

func main() {
	nums1 := []int{55, 30, 5, 4, 2}
	nums2 := []int{100, 20, 10, 10, 5}
	assertEq(2, maxDistance(nums1, nums2))

	nums1 = []int{2, 2, 2}
	nums2 = []int{10, 10, 1}
	assertEq(1, maxDistance(nums1, nums2))

	nums1 = []int{30, 29, 19, 5}
	nums2 = []int{25, 25, 25, 25, 25}
	assertEq(2, maxDistance(nums1, nums2))
}

func assertEq[T comparable](a, b T) {
	if a != b {
		fmt.Printf("Failed: %v != %v\n", a, b)
	} else {
		fmt.Printf("Ok: %v\n", a)
	}
}

func maxDistance(nums1 []int, nums2 []int) int {
	var i, j, res int
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] > nums2[j] {
			i++
			j = max(j, i)
			continue
		}
		res = max(res, j-i)
		j++
	}
	return res
}
