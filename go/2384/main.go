package main

import (
	"fmt"
	"slices"
)

func main() {
	nums := "444947137"
	assertEq("7449447", largestPalindromic(nums))

	nums = "00009"
	assertEq("9", largestPalindromic(nums))

	nums = "00001105"
	assertEq("1005001", largestPalindromic(nums))
}

func assertEq[T comparable](a, b T) {
	if a != b {
		fmt.Printf("Failed: %v != %v\n", a, b)
	} else {
		fmt.Printf("Ok: %v\n", a)
	}
}

func largestPalindromic(num string) string {
	count := make([]int, 10)
	for _, n := range []byte(num) {
		count[n-'0']++
	}

	var res []byte
	for i := 9; i >= 0; i-- {
		if i == 0 && len(res) == 0 {
			break
		}
		for count[i] >= 2 {
			res = append(res, byte(i)+'0')
			count[i] -= 2
		}
	}

	right := make([]byte, len(res))
	copy(right, res)
	slices.Reverse(right)

	i := 9
	for ; i >= 0; i-- {
		if count[i] > 0 {
			res = append(res, byte(i)+'0')
			break
		}
	}

	res = append(res, right...)
	return string(res)
}
