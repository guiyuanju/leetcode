package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxProfit([]int{1, 2, 3, 0, 2}))
	fmt.Println(maxProfit([]int{1}))
}

func maxProfit2(prices []int) int {
	var dp func(i int, holding bool, cooling bool) int
	dp = func(i int, holding bool, cooling bool) int {
		if i >= len(prices) {
			return 0
		}

		if cooling {
			return dp(i+1, holding, false)
		}

		res := dp(i+1, holding, false)
		if holding {
			res = max(res, prices[i]+dp(i+1, false, true))
		} else {
			res = max(res, -prices[i]+dp(i+1, true, false))
		}

		return res
	}

	return dp(0, false, false)
}

func maxProfit(prices []int) int {
	dp := make([][][]int, len(prices)+1)
	for i := range dp {
		// 1 ~ holding
		dp[i] = make([][]int, 2)
		for j := range dp[i] {
			// 1 ~ cooling down
			dp[i][j] = make([]int, 2)
		}
	}

	for i := len(prices) - 1; i >= 0; i-- {
		for holding := range 2 {
			for cooling := range 2 {
				if cooling == 1 {
					dp[i][holding][cooling] = dp[i+1][holding][0]
					continue
				}

				dp[i][holding][cooling] = dp[i+1][holding][0]
				if holding == 1 {
					dp[i][holding][cooling] = max(dp[i][holding][cooling], prices[i]+dp[i+1][0][1])
				} else {
					dp[i][holding][cooling] = max(dp[i][holding][cooling], -prices[i]+dp[i+1][1][0])
				}
			}
		}
	}

	return dp[0][0][0]
}
