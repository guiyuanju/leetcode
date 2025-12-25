package main

import "fmt"

func main() {
	fmt.Println(combine(4, 2))
	fmt.Println(combine(1, 1))
}

func combine(n int, k int) [][]int {
	var res [][]int
	var dp func(i int, cur []int)
	dp = func(i int, cur []int) {
		if len(cur) == k {
			tmp := make([]int, k)
			copy(tmp, cur)
			res = append(res, tmp)
			return
		}

		for j := i; j <= n; j++ {
			dp(j+1, append(cur, j))
		}
	}

	dp(1, nil)

	return res
}
