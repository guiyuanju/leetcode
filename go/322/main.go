package main

import (
	"math"

	"github.com/guiyuanju/lcutils/assert"
)

func main() {
	assert.Eq(3, coinChange([]int{1, 2, 5}, 11))
	assert.Eq(-1, coinChange([]int{2}, 3))
	assert.Eq(0, coinChange([]int{1}, 0))
}

func coinChange2(coins []int, amount int) int {
	var dp func(a int) int
	dp = func(a int) int {
		if a == 0 {
			return 0
		}
		if a < 0 {
			return -1
		}

		res := math.MaxInt
		for _, c := range coins {
			r := dp(a - c)
			if r != -1 {
				res = min(res, r)
			}
		}

		if res == math.MaxInt {
			return -1
		}
		return 1 + res
	}

	return dp(amount)
}

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		dp[i] = -1
	}

	for i := 1; i <= amount; i++ {
		for _, c := range coins {
			remain := i - c
			if remain < 0 {
				continue
			}
			if dp[remain] < 0 {
				continue
			}
			if dp[i] < 0 {
				dp[i] = dp[remain] + 1
				continue
			}
			dp[i] = min(dp[i], dp[remain]+1)
		}
	}

	return dp[amount]
}
