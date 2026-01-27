package main

import (
	"fmt"
	"reflect"
)

func main() {
	assertEq(4, findMaxForm([]string{"10", "0001", "111001", "1", "0"}, 5, 3))
	assertEq(2, findMaxForm([]string{"10", "0", "1"}, 1, 1))
	assertEq(3, findMaxForm([]string{"10", "0001", "111001", "1", "0"}, 4, 3))
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
	strs01 := make([][2]int, len(strs))
	for i, s := range strs {
		for _, c := range s {
			if c == '0' {
				strs01[i][0]++
			} else {
				strs01[i][1]++
			}
		}
	}

	memo := map[[3]int]int{}

	var dp func(i, m, n int) int
	dp = func(i, m, n int) int {
		if v, ok := memo[[3]int{i, m, n}]; ok {
			return v
		}

		if i >= len(strs) {
			return 0
		}

		res := dp(i+1, m, n)
		if m >= strs01[i][0] && n >= strs01[i][1] {
			res = max(1+dp(i+1, m-strs01[i][0], n-strs01[i][1]), res)
		}

		memo[[3]int{i, m, n}] = res

		return res
	}

	return dp(0, m, n)
}

func findMaxForm_bu(strs []string, m int, n int) int {
	strs01 := make([][2]int, len(strs))
	for i, s := range strs {
		for _, c := range s {
			if c == '0' {
				strs01[i][0]++
			} else {
				strs01[i][1]++
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

	for i := len(dp) - 2; i >= 0; i-- {
		for j := 0; j < len(dp[i]); j++ {
			for k := 0; k < len(dp[i][j]); k++ {
				dp[i][j][k] = dp[i+1][j][k]
				if j >= strs01[i][0] && k >= strs01[i][1] {
					dp[i][j][k] = max(dp[i][j][k], 1+dp[i+1][j-strs01[i][0]][k-strs01[i][1]])
				}
			}
		}
	}

	return dp[0][m][n]
}
