package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxProfit(2, []int{2, 4, 1}))
	fmt.Println(maxProfit(2, []int{3, 2, 6, 5, 0, 3}))
}

func maxProfit(k int, prices []int) int {
	// return maxProfit_td(k, prices)
	return maxProfit_bu(k, prices)
}

func maxProfit_td(k int, prices []int) int {
	type key struct {
		i, remain int
		holding   bool
	}
	memo := map[key]int{}
	var dp func(i, remain int, holding bool) int
	dp = func(i, remain int, holding bool) int {
		if v, ok := memo[key{i, remain, holding}]; ok {
			return v
		}

		if i == len(prices) {
			return 0
		}

		res := dp(i+1, remain, holding)
		if holding {
			res = max(prices[i]+dp(i+1, remain, false), res)
		} else if remain > 0 {
			res = max(-prices[i]+dp(i+1, remain-1, true), res)
		}

		memo[key{i, remain, holding}] = res

		return res
	}

	return dp(0, k, false)
}

func maxProfit_bu(k int, prices []int) int {
	dp := make([][][]int, len(prices)+1)
	for i := range dp {
		dp[i] = make([][]int, k+1)
		for j := range len(dp[i]) {
			dp[i][j] = make([]int, 2)
		}
	}

	for i := len(prices) - 1; i >= 0; i-- {
		for j := range k + 1 {
			for k := range 2 {
				dp[i][j][k] = dp[i+1][j][k]
				if k == 1 {
					dp[i][j][k] = max(prices[i]+dp[i+1][j][0], dp[i][j][k])
				} else if j > 0 {
					dp[i][j][k] = max(-prices[i]+dp[i+1][j-1][1], dp[i][j][k])
				}
			}
		}
	}

	return dp[0][k][0]
}
