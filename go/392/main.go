package main

import "fmt"

func main() {
	s := "abc"
	t := "ahbgdc"
	fmt.Println(isSubsequence(s, t))
	s = "axc"
	t = "ahbgdc"
	fmt.Println(isSubsequence(s, t))
	s = "b"
	t = "c"
	fmt.Println(isSubsequence(s, t))
}

func isSubsequence(s string, t string) bool {
	i := 0
	j := 0
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
			j++
		} else {
			j++
		}
	}
	return i >= len(s)
}
