package main

import "fmt"

func main() {
	sweetness := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	assertEq(6, maximizeSweetness(sweetness, 5))

	sweetness = []int{5, 6, 7, 8, 9, 1, 2, 3, 4}
	assertEq(1, maximizeSweetness(sweetness, 8))

	sweetness = []int{1, 2, 2, 1, 2, 2, 1, 2, 2}
	assertEq(5, maximizeSweetness(sweetness, 2))
}

func assertEq[T comparable](a, b T) {
	if a != b {
		fmt.Printf("Failed: %v != %v\n", a, b)
	} else {
		fmt.Printf("Ok: %v\n", a)
	}
}

func maximizeSweetness(sweetness []int, k int) int {
	check := func(g int) bool {
		var cur, count int
		for _, s := range sweetness {
			cur += s
			if cur >= g {
				cur = 0
				count++
			}
		}
		return count >= k+1
	}

	lo := 0
	hi := 1
	for _, s := range sweetness {
		hi += s
	}

	for lo < hi {
		mid := lo + (hi-lo)/2
		if check(mid) {
			lo = mid + 1
		} else {
			hi = mid
		}
	}

	return lo - 1
}
