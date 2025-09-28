package main

import "fmt"

func main() {
	fmt.Println(combine(4, 2))
	fmt.Println(combine(1, 1))
}

func combine(n int, k int) [][]int {
	var res [][]int
	var bt func(i int, cur []int)
	bt = func(i int, cur []int) {
		if len(cur) == k {
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			res = append(res, tmp)
			return
		}

		for j := i; j <= n; j++ {
			bt(j+1, append(cur, j))
		}
	}

	bt(1, nil)

	return res
}
