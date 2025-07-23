package main

import (
	"github.com/guiyuanju/lcutils/assert"
)

func main() {
	assert.Eq(2, climbStairs(2))
	assert.Eq(3, climbStairs(3))
}

func climbStairs2(n int) int {
	var dp func(i int) int
	dp = func(i int) int {
		if i <= 2 {
			return i
		}

		return dp(i-1) + dp(i-2)
	}

	return dp(n)
}

func climbStairs(n int) int {
	dp1 := 1
	dp2 := 2
	for i := 3; i <= n; i++ {
		dp1 = dp1 + dp2
		dp1, dp2 = dp2, dp1
	}
	if n == 1 {
		return dp1
	}
	return dp2
}
