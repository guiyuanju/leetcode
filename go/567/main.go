package main

import "fmt"

func checkInclusion(s1 string, s2 string) bool {
	count := map[byte]int{}
	for _, c := range s1 {
		count[byte(c)] += 1
	}

	var i, j, cur int
	for j < len(s2) {
		if c, ok := count[s2[j]]; ok {
			if c == 0 {
				for ; s2[i] != s2[j]; i += 1 {
					count[s2[i]] += 1
					cur -= 1
				}
				count[s2[i]] += 1
				cur -= 1
				i += 1
			} else {
				count[s2[j]] -= 1
				cur += 1
				if cur == len(s1) {
					return true
				}
				j += 1
			}
		} else {
			for i < j {
				count[s2[i]] += 1
				cur -= 1
				i += 1
			}
			i += 1
			j += 1
		}
	}
	return false
}

// 438
func findAnagrams(s string, p string) []int {
	count := map[byte]int{}
	for _, c := range p {
		count[byte(c)] += 1
	}

	var i, j, cur int
	var res []int
	for j < len(s) {
		if c, ok := count[s[j]]; ok {
			if c == 0 {
				for ; s[i] != s[j]; i += 1 {
					count[s[i]] += 1
					cur -= 1
				}
				count[s[i]] += 1
				cur -= 1
				i += 1
			} else {
				count[s[j]] -= 1
				cur += 1
				if cur == len(p) {
					res = append(res, i)
				}
				j += 1
			}
		} else {
			for i < j {
				count[s[i]] += 1
				cur -= 1
				i += 1
			}
			i += 1
			j += 1
		}
	}
	return res
}

func main() {
	s1 := "ab"
	s2 := "eidbaooo"
	fmt.Println(checkInclusion(s1, s2))
	s1 = "ab"
	s2 = "eidboaoo"
	fmt.Println(checkInclusion(s1, s2))
	s1 = "abbc"
	s2 = "cbabadcbbabbcbabaabccbabc"
	fmt.Println(checkInclusion(s1, s2))
}
