package main

import "fmt"

func main() {
	s := "aab"
	fmt.Println(partition(s))

	s = "a"
	fmt.Println(partition(s))
}

func partition(s string) [][]string {
	var res [][]string
	var bt func(cur []string, i int)
	bt = func(cur []string, i int) {
		if i == len(s) {
			res = append(res, append([]string(nil), cur...))
			return
		}

		for j := i; j < len(s); j++ {
			if isPanlindrome(s[i : j+1]) {
				bt(append(cur, s[i:j+1]), j+1)
			}
		}
	}

	bt(nil, 0)

	return res
}

func isPanlindrome(s string) bool {
	i := 0
	j := len(s) - 1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
