package main

import "fmt"

func main() {
	fmt.Println(maxValueOfCoins([][]int{{1, 100, 3}, {7, 8, 9}}, 2))
	fmt.Println(maxValueOfCoins([][]int{{100}, {100}, {100}, {100}, {100}, {100}, {1, 1, 1, 1, 1, 1, 700}}, 7))
}

func toString(xs []int, x int) string {
	return fmt.Sprintf("%v%d", xs, x)
}

// top down
func maxValueOfCoins2(piles [][]int, k int) int {
	memo := map[[2]int]int{}

	var dp func(i int, k int) int
	dp = func(i int, k int) int {
		if v, ok := memo[[2]int{i, k}]; ok {
			return v
		}

		if k <= 0 || i < 0 {
			return 0
		}

		var sum int
		res := dp(i-1, k)
		for j := 1; j <= k && j <= len(piles[i]); j++ {
			sum += piles[i][j-1]
			res = max(res, sum+dp(i-1, k-j))
		}

		memo[[2]int{i, k}] = res
		return res
	}

	return dp(len(piles)-1, k)
}

// bottom up
func maxValueOfCoins(piles [][]int, k int) int {
	dp := make([][]int, k+1)
	for i := range dp {
		dp[i] = make([]int, len(piles)+1)
	}

	for i := 1; i <= k; i++ {
		for j := 1; j <= len(piles); j++ {
			var sum int
			res := dp[i][j-1]
			for k := 1; k <= i && k <= len(piles[j-1]); k++ {
				sum += piles[j-1][k-1]
				res = max(res, sum+dp[i-k][j-1])
			}
			dp[i][j] = res
		}
	}

	return dp[k][len(piles)]
}
