package main

import (
	"fmt"
)

func main() {
	fmt.Println(numRollsToTarget(1, 6, 3))
	fmt.Println(numRollsToTarget(2, 6, 7))
	fmt.Println(numRollsToTarget(30, 30, 500))
}

func numRollsToTargetBT(n int, k int, target int) int {
	var res int
	var bt func(cur []int, i, remian int)
	bt = func(cur []int, i, remain int) {
		if i > n {
			if remain == 0 {
				res++
			}
			return
		}

		for v := 1; v <= k && v < remain; v++ {
			bt(append(cur, v), i+1, remain-v)
		}
	}

	bt(nil, 1, target)

	return res
}

func numRollsToTargetTD(n int, k int, target int) int {
	memo := map[[2]int]int{}

	var dp func(i, remian int) int
	dp = func(i, remain int) int {
		if v, ok := memo[[2]int{i, remain}]; ok {
			return v
		}

		if i == 1 {
			if remain <= k {
				return 1
			}
			return 0
		}

		var res int
		for j := 1; j <= k && j < remain; j++ {
			res += dp(i-1, remain-j)
			res %= (1e9 + 7)
		}

		memo[[2]int{i, remain}] = res
		return res
	}

	return dp(n, target)
}

func numRollsToTarget(n int, k int, target int) int {
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, target+1)
	}
	for i := range target + 1 {
		if i <= k {
			dp[1][i] = 1
		}
	}

	for i := 2; i <= n; i++ {
		for j := 0; j <= target; j++ {
			for v := 1; v <= k && v < j; v++ {
				dp[i][j] += dp[i-1][j-v]
				dp[i][j] %= 1e9 + 7
			}
		}
	}

	return dp[n][target]
}
