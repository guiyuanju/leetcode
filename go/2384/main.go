package main

import (
	"fmt"
	"maps"
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
	count := map[byte]int{}
	for _, n := range num {
		count[byte(n)]++
	}

	descend := slices.Collect(maps.Keys(count))
	slices.SortFunc(descend, func(a, b byte) int {
		return int(b) - int(a)
	})

	var res []byte
	for _, b := range descend {
		if len(res) == 0 && b == '0' {
			continue
		}
		for count[b] >= 2 {
			res = append(res, b)
			count[b] -= 2
		}
	}
	length := len(res)

	for _, b := range descend {
		if count[b] > 0 {
			res = append(res, b)
			break
		}
	}

	for i := range length {
		res = append(res, res[length-i-1])
	}

	return string(res)
}
