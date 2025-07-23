package main

import (
	"github.com/guiyuanju/lcutils/assert"
)

func main() {
	assert.Eq(int64(5), mostPoints([][]int{{3, 2}, {4, 3}, {4, 4}, {2, 5}}))
	assert.Eq(int64(7), mostPoints([][]int{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}}))
}

func mostPoints2(questions [][]int) int64 {
	var dp func(i int) int
	dp = func(i int) int {
		if i >= len(questions) {
			return 0
		}
		return max(questions[i][0]+dp(i+questions[i][1]+1), dp(i+1))
	}

	return int64(dp(0))
}

func mostPoints(questions [][]int) int64 {
	dp := make([]int, len(questions))
	dp[len(dp)-1] = questions[len(questions)-1][0]
	for i := len(dp) - 2; i >= 0; i-- {
		var tmp int
		if i+questions[i][1]+1 < len(dp) {
			tmp = dp[i+questions[i][1]+1]
		}
		dp[i] = max(questions[i][0]+tmp, dp[i+1])
	}
	return int64(dp[0])
}
