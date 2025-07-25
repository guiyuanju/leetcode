package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{1, 3, 2, 8, 4, 9}, 2))
	fmt.Println(maxProfit([]int{1, 3, 7, 5, 10, 3}, 3))
}

func maxProfit2(prices []int, fee int) int {
	var dp func(i int, holding bool) int
	dp = func(i int, holding bool) int {
		if i >= len(prices) {
			return 0
		}

		res := dp(i+1, holding)
		if holding {
			res = max(res, prices[i]+dp(i+1, false))
		} else {
			res = max(res, -prices[i]-fee+dp(i+1, true))
		}
		return res
	}

	return dp(0, false)
}

func maxProfit(prices []int, fee int) int {
	dp := make([][]int, len(prices)+1)
	for i := range dp {
		// 1 ~ holding, 0 ~ not holding
		dp[i] = make([]int, 2)
	}

	for i := len(prices) - 1; i >= 0; i-- {
		for j := range 2 {
			dp[i][j] = dp[i+1][j]
			if j == 1 {
				dp[i][j] = max(dp[i][j], prices[i]+dp[i+1][0])
			} else {
				dp[i][j] = max(dp[i][j], -prices[i]-fee+dp[i+1][1])
			}
		}
	}

	return dp[0][0]
}
