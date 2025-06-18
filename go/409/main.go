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
	counts := map[byte]int{}
	for _, r := range s {
		b := byte(r)
		counts[b]++
	}

	var hasSingle bool
	var count int
	for _, c := range counts {
		if c%2 > 0 {
			hasSingle = true
		}
		count += c / 2 * 2
	}

	if hasSingle {
		count++
	}

	return count
}
