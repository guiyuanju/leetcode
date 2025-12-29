package main

import (
	"fmt"
	"reflect"
)

func main() {
	assertEq(4, findMaxForm([]string{"10", "0001", "111001", "1", "0"}, 5, 3))
	assertEq(2, findMaxForm([]string{"10", "0", "1"}, 1, 1))
}

func assertEq(a, b any) {
	if reflect.DeepEqual(a, b) {
		fmt.Printf("Ok: %v\n", a)
	} else {
		fmt.Printf("Failed: %v != %v\n", a, b)
	}
}

func findMaxForm(strs []string, m int, n int) int {
	// return findMaxForm_td(strs, m, n)
	return findMaxForm_bu(strs, m, n)
}

func findMaxForm_td(strs []string, m int, n int) int {
	count := make([][2]int, len(strs))
	for i, s := range strs {
		for _, c := range s {
			if c == '0' {
				count[i][0]++
			} else {
				count[i][1]++
			}
		}
	}

	var dp func(i, m, n int) int
	dp = func(i, m, n int) int {
		if i >= len(strs) || (m == 0 && n == 0) {
			return 0
		}

		res := dp(i+1, m, n)
		if count[i][0] <= m && count[i][1] <= n {
			res = max(res, 1+dp(i+1, m-count[i][0], n-count[i][1]))
		}

		return res
	}

	return dp(0, m, n)
}

func findMaxForm_bu(strs []string, m int, n int) int {
	count := make([][2]int, len(strs))
	for i, s := range strs {
		for _, c := range s {
			if c == '0' {
				count[i][0]++
			} else {
				count[i][1]++
			}
		}
	}

	dp := make([][][]int, len(strs)+1)
	for i := range dp {
		dp[i] = make([][]int, m+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, n+1)
		}
	}

	for i := len(strs) - 1; i >= 0; i-- {
		for j := range m + 1 {
			for k := range n + 1 {
				dp[i][j][k] = dp[i+1][j][k]
				if count[i][0] <= j && count[i][1] <= k {
					dp[i][j][k] = max(dp[i][j][k], 1+dp[i+1][j-count[i][0]][k-count[i][1]])
				}
			}
		}
	}

	return dp[0][m][n]
}
