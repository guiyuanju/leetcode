package main

import "fmt"

func main() {
	fmt.Println(findAnagrams("cbaebabacd", "abc"))
	fmt.Println(findAnagrams("abab", "ab"))
}

func findAnagrams(s string, p string) []int {
	count := map[byte]int{}
	for _, c := range p {
		count[byte(c)]++
	}

	var res []int
	for i, c := range []byte(s) {
		count[c]--
		if i >= len(p) {
			count[s[i-len(p)]]++
		}
		if i >= len(p)-1 {
			var zeros int
			for _, v := range count {
				if v != 0 {
					break
				}
				zeros++
			}
			if zeros == len(count) {
				res = append(res, i-len(p)+1)
			}
		}
	}
	return res
}
