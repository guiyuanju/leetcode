package main

import "fmt"

func main() {
	fmt.Println(generate(5))
	fmt.Println(generate(1))
}

func generate(numRows int) [][]int {
	res := make([][]int, numRows)
	res[0] = []int{1}
	for i := 1; i < numRows; i++ {
		cur := make([]int, i+1)
		cur[0] = 1
		cur[i] = 1
		for j := 1; j < i; j++ {
			cur[j] = res[i-1][j-1] + res[i-1][j]
		}
		res[i] = cur
	}
	return res
}
