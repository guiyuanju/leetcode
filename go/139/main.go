package main

import "fmt"

func main() {
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
	fmt.Println(wordBreak("applepenapple", []string{"apple", "pen"}))
	fmt.Println(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
}

func wordBreak2(s string, wordDict []string) bool {
	m := make(map[string]bool, len(wordDict))
	for _, w := range wordDict {
		m[w] = true
	}

	var dp func(i, j int) bool
	dp = func(i, j int) bool {
		if i >= len(s) {
			return true
		}
		if j > len(s) {
			return false
		}

		return m[s[i:j]] && dp(j, j+1) || dp(i, j+1)
	}

	return dp(0, 1)
}

func wordBreak(s string, wordDict []string) bool {
	m := make(map[string]bool, len(wordDict))
	for _, w := range wordDict {
		m[w] = true
	}

	dp := make([][]bool, len(s)+1)
	for i := range dp {
		dp[i] = make([]bool, len(s)+2)
	}
	for i := range len(s) + 2 {
		dp[len(s)][i] = true
	}

	for i := len(s) - 1; i >= 0; i-- {
		for j := len(s); j > i; j-- {
			dp[i][j] = m[s[i:j]] && dp[j][j+1] || dp[i][j+1]
		}
	}

	return dp[0][1]
}
