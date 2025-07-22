package main

import "fmt"

func main() {
	cost := []int{10, 15, 20}
	fmt.Println(minCostClimbingStairs2(cost))

	cost = []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	fmt.Println(minCostClimbingStairs2(cost))
}

func minCostClimbingStairs(cost []int) int {
	memo := make([]int, len(cost)+1)
	for i := range memo {
		memo[i] = -1
	}
	var dp func(i int) int
	dp = func(i int) int {
		if i <= 1 {
			return 0
		}
		if memo[i] != -1 {
			return memo[i]
		}
		memo[i] = min(dp(i-1)+cost[i-1], dp(i-2)+cost[i-2])
		return memo[i]
	}

	return dp(len(cost))
}

func minCostClimbingStairs2(cost []int) int {
	dp := make([]int, len(cost)+1)
	dp[0] = 0
	dp[1] = 0

	for i := 2; i <= len(cost); i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}

	return dp[len(cost)]
}
