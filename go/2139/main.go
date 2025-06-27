package main

import "fmt"

func main() {
	assertEq(4, minMoves(5, 0))
	assertEq(7, minMoves(19, 2))
	assertEq(4, minMoves(10, 4))
	assertEq(999999996, minMoves(999999997, 0))
}

func assertEq[T comparable](a, b T) {
	if a != b {
		fmt.Printf("Failed: %v != %v\n", a, b)
	} else {
		fmt.Printf("Ok: %v\n", a)
	}
}

func minMoves(target int, maxDoubles int) int {
	var res int
	for target > 1 {
		if maxDoubles == 0 {
			return res + target - 1
		}
		if target%2 > 0 {
			target--
			res++
		}
		target /= 2
		maxDoubles--
		res++
	}
	return res
}
