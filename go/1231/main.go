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
	check := func(sweet int) bool {
		var i, j int
		for ; i <= k; i++ {
			var cur int
			for ; j < len(sweetness) && cur < sweet; j++ {
				cur += sweetness[j]
			}

			if cur < sweet {
				return false
			}

			if j == len(sweetness) && i < k {
				return false
			}
		}

		return true
	}

	left := 1
	right := 0
	for _, s := range sweetness {
		right += s
	}

	for left <= right {
		mid := left + (right-left)/2
		if check(mid) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return right
}
