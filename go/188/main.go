package main

import (
	"fmt"
	"reflect"
)

func main() {
	assertEq(2, maxProfit(2, []int{2, 4, 1}))
	assertEq(7, maxProfit(2, []int{3, 2, 6, 5, 0, 3}))
}

func assertEq(a, b any) {
	if reflect.DeepEqual(a, b) {
		fmt.Printf("Ok: %v\n", a)
	} else {
		fmt.Printf("Failed: %v != %v\n", a, b)
	}
}

func maxProfit(k int, prices []int) int {
	// return maxProfit_td(k, prices)
	return maxProfit_bu(k, prices)
}

func maxProfit_td(k int, prices []int) int {
	memo := map[[3]int]int{}

	var dp func(i int, k int, hold int) int
	dp = func(i int, k int, hold int) int {
		if v, ok := memo[[3]int{i, k, hold}]; ok {
			return v
		}

		if i >= len(prices) {
			return 0
		}

		res := dp(i+1, k, hold)
		if k > 0 && hold == 0 {
			res = max(res, -prices[i]+dp(i+1, k-1, 1))
		}
		if hold == 1 {
			res = max(res, prices[i]+dp(i+1, k, 0))
		}

		memo[[3]int{i, k, hold}] = res

		return res
	}

	return dp(0, k, 0)
}

func maxProfit_bu(k int, prices []int) int {
	dp := make([][][]int, len(prices)+1)
	for i := range dp {
		dp[i] = make([][]int, k+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, 2)
		}
	}

	for i := len(prices) - 1; i >= 0; i-- {
		for j := range k + 1 {
			for hold := range 2 {
				dp[i][j][hold] = dp[i+1][j][hold]
				if j > 0 && hold == 0 {
					dp[i][j][hold] = max(dp[i][j][hold], -prices[i]+dp[i+1][j-1][1])
				}
				if hold == 1 {
					dp[i][j][hold] = max(dp[i][j][hold], prices[i]+dp[i+1][j][0])
				}
			}
		}
	}

	return dp[0][k][0]
}
