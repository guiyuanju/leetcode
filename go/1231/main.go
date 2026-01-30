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
	check := func(s int) bool {
		var cur, count int
		for _, v := range sweetness {
			cur += v
			if cur >= s {
				cur = 0
				count++
			}
		}
		return count >= k+1
	}

	var i, j int
	for _, s := range sweetness {
		j += s
	}
	j++

	for i < j {
		mid := i + (j-i)/2
		if check(mid) {
			i = mid + 1
		} else {
			j = mid
		}
	}

	return i - 1
}
