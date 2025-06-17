package main

import (
	"fmt"
	"strconv"
)

func main() {
	assertEq(9969, maximum69Number(9669))
	assertEq(9999, maximum69Number(9996))
	assertEq(9999, maximum69Number(9999))
}

func assertEq[T comparable](a, b T) {
	if a != b {
		fmt.Printf("Failed: %v != %v\n", a, b)
	} else {
		fmt.Printf("Ok: %v\n", a)
	}
}

func maximum69Number(num int) int {
	s := []byte(strconv.Itoa(num))
	for i, v := range s {
		if v == '6' {
			s[i] = '9'
			break
		}
	}
	res, _ := strconv.Atoi(string(s))
	return res
}
