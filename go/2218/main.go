package main

import (
	"fmt"
	"reflect"
)

func main() {
	assertEq(101, maxValueOfCoins([][]int{{1, 100, 3}, {7, 8, 9}}, 2))
	assertEq(706, maxValueOfCoins([][]int{{100}, {100}, {100}, {100}, {100}, {100}, {1, 1, 1, 1, 1, 1, 700}}, 7))
	assertEq(494, maxValueOfCoins([][]int{{37, 88}, {51, 64, 65, 20, 95, 30, 26}, {9, 62, 20}, {44}}, 9))
}

func assertEq(a, b any) {
	if reflect.DeepEqual(a, b) {
		fmt.Printf("Ok: %v\n", a)
	} else {
		fmt.Printf("Failed: %v != %v\n", a, b)
	}
}

func maxValueOfCoins(piles [][]int, k int) int {
	return maxValueOfCoins_td(piles, k)
	// return maxValueOfCoins_bu(piles, k)
}

func maxValueOfCoins_td(piles [][]int, k int) int {
	var dp func(i int, k int) int
	dp = func(i int, k int) int {
		if i >= len(piles) || k == 0 {
			return 0
		}

		var cur int
		res := dp(i+1, k)
		for j := 0; j < len(piles[i]) && j < k; j++ {
			cur += piles[i][j]
			res = max(res, cur+dp(i+1, k-j-1))
		}

		return res
	}

	return dp(0, k)
}

func maxValueOfCoins_bu(piles [][]int, k int) int {
	dp := make([][]int, len(piles)+1)
	for i := range dp {
		dp[i] = make([]int, k+1)
	}

	for i := len(piles) - 1; i >= 0; i-- {
		for j := 1; j <= k; j++ {
			var cur int
			dp[i][j] = dp[i+1][j]
			for k := range min(j, len(piles[i])) {
				cur += piles[i][k]
				dp[i][j] = max(dp[i][j], cur+dp[i+1][j-k-1])
			}
		}
	}

	return dp[0][k]
}
