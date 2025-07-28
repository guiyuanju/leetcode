package main

import "fmt"

func main() {
	fmt.Println(tribonacci(4))
	fmt.Println(tribonacci(25))
}

func tribonacci(n int) int {
	t0, t1, t2 := 0, 1, 1
	if n == 0 {
		return 0
	}
	if n <= 2 {
		return 1
	}
	for range n - 2 {
		t3 := t0 + t1 + t2
		t0, t1, t2 = t1, t2, t3
	}
	return t2
}
