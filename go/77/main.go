package main

import "fmt"

func main() {
	fmt.Println(combine(4, 2))
	fmt.Println(combine(1, 1))
}

func combine(n int, k int) [][]int {
	var res [][]int
	var bt func(cur []int, i int)
	bt = func(cur []int, i int) {
		if len(cur) == k {
			res = append(res, append([]int(nil), cur...))
			return
		}
		for j := i; j <= n; j++ {
			cur = append(cur, j)
			bt(cur, j+1)
			cur = cur[:len(cur)-1]
		}
	}
	bt(nil, 1)
	return res
}
