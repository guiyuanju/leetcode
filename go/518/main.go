package main

import "fmt"

func main() {
	fmt.Println(change(5, []int{1, 2, 5}))
	fmt.Println(change(3, []int{2}))
	fmt.Println(change(10, []int{10}))
}

func change2(amount int, coins []int) int {
	var dp func(i int, remain int) int
	dp = func(i int, remain int) int {
		if remain == 0 {
			return 1
		}
		if i >= len(coins) {
			return 0
		}

		var res int
		for j := range remain/coins[i] + 1 {
			res += dp(i+1, remain-j*coins[i])
		}
		return res
	}

	return dp(0, amount)
}

func change(amount int, coins []int) int {
	dp := make([][]int, len(coins)+1)
	for i := range dp {
		dp[i] = make([]int, amount+1)
		dp[i][0] = 1
	}

	for i := len(coins) - 1; i >= 0; i-- {
		for j := 1; j <= amount; j++ {
			for k := range j/coins[i] + 1 {
				dp[i][j] += dp[i+1][j-k*coins[i]]
			}
		}
	}

	return dp[0][amount]
}
