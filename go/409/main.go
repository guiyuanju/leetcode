package main

import "fmt"

func main() {
	s := "abccccdd"
	assertEq(7, longestPalindrome(s))

	s = "a"
	assertEq(1, longestPalindrome(s))

	s = "ccc"
	assertEq(3, longestPalindrome(s))
}

func assertEq[T comparable](a, b T) {
	if a != b {
		fmt.Printf("Failed: %v != %v\n", a, b)
	} else {
		fmt.Printf("Ok: %v\n", a)
	}
}

func longestPalindrome(s string) int {
	count := map[byte]int{}
	for _, c := range []byte(s) {
		count[c]++
	}
	var res int
	var hasOdd bool
	for _, v := range count {
		if v%2 == 1 {
			hasOdd = true
		}
		res += v / 2 * 2
	}
	if hasOdd {
		return res + 1
	}
	return res
}
