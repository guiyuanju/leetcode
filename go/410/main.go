package main

import "fmt"

func main() {
	nums := []int{7, 2, 5, 10, 8}
	assertEq(18, splitArray(nums, 2))

	nums = []int{1, 2, 3, 4, 5}
	assertEq(9, splitArray(nums, 2))
}

func assertEq[T comparable](a, b T) {
	if a != b {
		fmt.Printf("Failed: %v != %v\n", a, b)
	} else {
		fmt.Printf("Ok: %v\n", a)
	}
}

func splitArray(nums []int, k int) int {
	check := func(g int) bool {
		var cur, count int
		for _, n := range nums {
			cur += n
			if cur > g {
				cur = n
				count++
			}
		}
		return count+1 <= k
	}

	var lo, hi int
	for _, n := range nums {
		hi += n
		lo = max(lo, n)
	}
	hi++

	for lo < hi {
		mid := lo + (hi-lo)/2
		if check(mid) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}

	return lo
}
