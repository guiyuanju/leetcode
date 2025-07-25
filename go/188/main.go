package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxProfit(2, []int{2, 4, 1}))
	fmt.Println(maxProfit(2, []int{3, 2, 6, 5, 0, 3}))
}

// top down (memo omiteed)
func maxProfit2(k int, prices []int) int {
	var dp func(i int, holding bool, remain int) int
	dp = func(i int, holding bool, remain int) int {
		if i >= len(prices) || remain <= 0 {
			return 0
		}

		res := dp(i+1, holding, remain)
		if !holding {
			res = max(res, -prices[i]+dp(i+1, true, remain))
		} else {
			res = max(res, prices[i]+dp(i+1, false, remain-1))
		}
		return res
	}

	return dp(0, false, k)
}

// bottom up
func maxProfit(k int, prices []int) int {
	dp := make([][]map[bool]int, k+1)
	for i := range k + 1 {
		dp[i] = make([]map[bool]int, len(prices)+1)
		for j := range dp[i] {
			dp[i][j] = map[bool]int{}
		}
	}

	for i := 1; i < k+1; i++ {
		for j := len(prices) - 1; j >= 0; j-- {
			cur := dp[i][j]
			cur[false] = max(dp[i][j+1][false], -prices[j]+dp[i][j+1][true])
			cur[true] = max(dp[i][j+1][true], prices[j]+dp[i-1][j+1][false])
		}
	}

	return dp[k][0][false]
}
