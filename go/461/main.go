package main

import "fmt"

func main() {
	fmt.Println(hammingDistance(1, 4))
	fmt.Println(hammingDistance(3, 1))
}

func hammingDistance(x int, y int) int {
	z := x ^ y
	var res int
	for i := range 32 {
		res += (z >> i) & 0x1
	}
	return res
}
