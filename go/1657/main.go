package main

import "fmt"

func main() {
	fmt.Println(closeStrings("abc", "bca"))
	fmt.Println(closeStrings("a", "aa"))
	fmt.Println(closeStrings("cabbba", "abbccc"))
	fmt.Println(closeStrings("uau", "ssx"))
}

func closeStrings(word1 string, word2 string) bool {
	count1 := map[byte]int{}
	count2 := map[byte]int{}
	for _, c := range []byte(word1) {
		count1[c]++
	}
	for _, c := range []byte(word2) {
		count2[c]++
	}
	if len(count1) != len(count2) {
		return false
	}
	for c := range count1 {
		if count2[c] == 0 {
			return false
		}
	}

	freq1 := map[int]int{}
	for _, c := range count1 {
		freq1[c]++
	}
	freq2 := map[int]int{}
	for _, c := range count2 {
		freq2[c]++
	}
	if len(freq1) != len(freq2) {
		return false
	}
	for k, v := range freq1 {
		if freq2[k] != v {
			return false
		}
	}

	return true
}
