package main

import (
	"fmt"
)

func main() {
	fmt.Println(numRollsToTarget(1, 6, 3))
	fmt.Println(numRollsToTarget(2, 6, 7))
	fmt.Println(numRollsToTarget(30, 30, 500))
}

func numRollsToTarget(n int, k int, target int) int {
	// return numRollsToTargetTD(n, k, target)
	return numRollsToTargetBU(n, k, target)
}

func numRollsToTargetTD(n int, k int, target int) int {
	memo := map[[2]int]int{}

	var dp func(i int, target int) int
	dp = func(i int, target int) int {
		if v, ok := memo[[2]int{i, target}]; ok {
			return v
		}

		if i == n {
			if target == 0 {
				return 1
			}
			return 0
		}

		var res int
		for j := 1; j <= k; j++ {
			if target-j < 0 {
				break
			}
			res += dp(i+1, target-j)
			res %= 1e9 + 7
		}

		memo[[2]int{i, target}] = res

		return res
	}

	return dp(0, target)
}

func numRollsToTargetBU(n int, k int, target int) int {
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, target+1)
	}
	dp[n][0] = 1

	for i := n - 1; i >= 0; i-- {
		for remain := 1; remain <= target; remain++ {
			for j := 1; j <= k; j++ {
				if remain-j < 0 {
					break
				}
				dp[i][remain] += dp[i+1][remain-j]
				dp[i][remain] %= 1e9 + 7
			}
		}
	}

	return dp[0][target]
}
