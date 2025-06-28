package main

import (
	"fmt"
	"math"
)

func main() {
	assertEq(4, makeIntegerBeautiful(16, 6))
	assertEq(33, makeIntegerBeautiful(467, 6))
	assertEq(0, makeIntegerBeautiful(1, 1))
}

func assertEq[T comparable](a, b T) {
	if a != b {
		fmt.Printf("Failed: %v != %v\n", a, b)
	} else {
		fmt.Printf("Ok: %v\n", a)
	}
}

func makeIntegerBeautiful(n int64, target int) int64 {
	digits := []int{}
	for n > 0 {
		digits = append(digits, int(n%10))
		n /= 10
	}

	var sum int
	for _, d := range digits {
		sum += d
	}

	if sum <= target {
		return 0
	}

	var i int
	var res int64
	var carry bool
	for sum > target-1 && i < len(digits) {
		sum -= digits[i]
		left := int64(10 - digits[i])
		if carry {
			left--
		}
		carry = true
		res += left * int64(math.Pow10(i))
		i++
	}

	return res
}
