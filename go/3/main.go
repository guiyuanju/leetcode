package main

import "fmt"

func main() {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s))
	s = "bbbbb"
	fmt.Println(lengthOfLongestSubstring(s))
	s = "pwwkew"
	fmt.Println(lengthOfLongestSubstring(s))
	s = " "
	fmt.Println(lengthOfLongestSubstring(s))
	s = "aab"
	fmt.Println(lengthOfLongestSubstring(s))
	s = "abba"
	fmt.Println(lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) int {
	count := map[byte]int{}
	var i, j, maxLen int
	for ; j < len(s); j++ {
		count[s[j]]++
		for count[s[j]] > 1 {
			count[s[i]]--
			i++
		}
		maxLen = max(maxLen, j-i+1)
	}
	return maxLen
}
func lengthOfLongestSubstring2(s string) int {
	mIdx := map[byte]int{}
	var i, j, maxLen int
	for ; j < len(s); j++ {
		if idx, ok := mIdx[s[j]]; ok {
			// 防止重复的char在窗口左边缘之前, 导致i往前跳
			i = max(i, idx+1)
		}
		mIdx[s[j]] = j
		maxLen = max(maxLen, j-i+1)
	}
	return maxLen
}
