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
	var dp func(i int, k int, holding int) int
	dp = func(i int, k int, holding int) int {
		if i >= len(prices) {
			return 0
		}

		res := dp(i+1, k, holding)
		if holding == 1 {
			res = max(res, prices[i]+dp(i+1, k, 0))
		} else if k > 0 {
			res = max(res, -prices[i]+dp(i+1, k-1, 1))
		}

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
		for curK := 0; curK <= k; curK++ {
			copy(dp[i][curK], dp[i+1][curK])
			dp[i][curK][1] = max(dp[i][curK][1], prices[i]+dp[i+1][curK][0])
			if curK > 0 {
				dp[i][curK][0] = max(dp[i][curK][0], -prices[i]+dp[i+1][curK-1][1])
			}
		}
	}

	return dp[0][k][0]
}
