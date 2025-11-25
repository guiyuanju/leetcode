package main

import "fmt"

func main() {
	fmt.Println(findMaxForm([]string{"10", "0001", "111001", "1", "0"}, 5, 3))
	fmt.Println(findMaxForm([]string{"10", "0", "1"}, 1, 1))
}

func findMaxForm(strs []string, m int, n int) int {
	// return findMaxFormTD(strs, m, n)
	return findMaxFormBU(strs, m, n)
}

func findMaxFormTD(strs []string, m int, n int) int {
	count := make([][2]int, len(strs))
	for i, s := range strs {
		for _, c := range s {
			count[i][int(c-'0')]++
		}
	}

	memo := map[[3]int]int{}

	var dp func(i, m, n int) int
	dp = func(i, m, n int) int {
		if v, ok := memo[[3]int{i, m, n}]; ok {
			return v
		}

		if i >= len(strs) || (m == 0 && n == 0) {
			return 0
		}

		var res int
		if count[i][0] <= m && count[i][1] <= n {
			res = 1 + dp(i+1, m-count[i][0], n-count[i][1])
		}
		res = max(res, dp(i+1, m, n))

		memo[[3]int{i, m, n}] = res

		return res
	}

	return dp(0, m, n)
}

func findMaxFormBU(strs []string, m int, n int) int {
	dp := make([][][]int, len(strs)+1)
	for i := range dp {
		dp[i] = make([][]int, m+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, n+1)
		}
	}

	count := make([][2]int, len(strs))
	for i, s := range strs {
		for _, c := range s {
			count[i][int(c-'0')]++
		}
	}

	for i := len(strs) - 1; i >= 0; i-- {
		for j := 0; j <= m; j++ {
			for k := 0; k <= n; k++ {
				if count[i][0] <= j && count[i][1] <= k {
					dp[i][j][k] = 1 + dp[i+1][j-count[i][0]][k-count[i][1]]
				}
				dp[i][j][k] = max(dp[i][j][k], dp[i+1][j][k])
			}
		}
	}

	return dp[0][m][n]
}
