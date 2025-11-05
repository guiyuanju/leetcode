package main

import "fmt"

func main() {
	fmt.Println(maxValueOfCoins([][]int{{1, 100, 3}, {7, 8, 9}}, 2))
	fmt.Println(maxValueOfCoins([][]int{{100}, {100}, {100}, {100}, {100}, {100}, {1, 1, 1, 1, 1, 1, 700}}, 7))
}

func maxValueOfCoins(piles [][]int, k int) int {
	// return maxValueOfCoinsTD(piles, k)
	return maxValueOfCoinsBU(piles, k)
}

// top down
func maxValueOfCoinsTD(piles [][]int, k int) int {
	type key struct{ i, left int }
	memo := map[key]int{}

	var dp func(i, left int) int
	dp = func(i, left int) int {
		if v, ok := memo[key{i, left}]; ok {
			return v
		}

		if i == len(piles) || left == 0 {
			return 0
		}

		var cur int
		res := dp(i+1, left)
		for j := range min(left, len(piles[i])) {
			cur += piles[i][j]
			res = max(res, cur+dp(i+1, left-j-1))
		}

		memo[key{i, left}] = res

		return res
	}

	return dp(0, k)
}

func maxValueOfCoinsBU(piles [][]int, k int) int {
	dp := make([][]int, len(piles)+1)
	for i := range dp {
		dp[i] = make([]int, k+1)
	}

	for i := len(piles) - 1; i >= 0; i-- {
		for j := 1; j <= k; j++ {
			var cur int
			res := dp[i+1][j]
			for k := range min(j, len(piles[i])) {
				cur += piles[i][k]
				res = max(res, cur+dp[i+1][j-k-1])
			}
			dp[i][j] = res
		}
	}

	return dp[0][k]
}
