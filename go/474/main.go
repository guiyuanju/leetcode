package main

import "fmt"

func main() {
	fmt.Println(findMaxForm([]string{"10", "0001", "111001", "1", "0"}, 5, 3))
	fmt.Println(findMaxForm([]string{"10", "0", "1"}, 1, 1))
}

func findMaxForm(strs []string, m int, n int) int {
	mnlen := func(i int) (int, int) {
		var m, n int
		for _, b := range strs[i] {
			if byte(b) == '0' {
				m++
			} else {
				n++
			}
		}
		return m, n
	}

	memo := map[[3]int]int{}

	var dp func(i, m, n int) int
	dp = func(i, m, n int) int {
		if v, ok := memo[[3]int{i, m, n}]; ok {
			return v
		}

		if i < 0 || (m == 0 && n == 0) {
			return 0
		}

		res := dp(i-1, m, n)
		a, b := mnlen(i)
		if a <= m && b <= n {
			res = max(res, 1+dp(i-1, m-a, n-b))
		}

		memo[[3]int{i, m, n}] = res

		return res
	}

	return dp(len(strs)-1, m, n)
}
