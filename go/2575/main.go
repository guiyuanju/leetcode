package main

import "fmt"

func main() {
	fmt.Println(divisibilityArray("998244353", 3))
	fmt.Println(divisibilityArray("1010", 10))
}

func divisibilityArray(word string, m int) []int {
	res := make([]int, len(word))
	var cur int
	for i, c := range word {
		d := int(byte(c) - '0')
		cur = cur*10 + d
		cur %= m
		if cur == 0 {
			res[i] = 1
		}
	}
	return res
}
