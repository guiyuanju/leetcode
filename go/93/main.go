package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "25525511135"
	fmt.Println(restoreIpAddresses(s))

	s = "0000"
	fmt.Println(restoreIpAddresses(s))

	s = "101023"
	fmt.Println(restoreIpAddresses(s))

	s = "0"
	fmt.Println(restoreIpAddresses(s))
}

func restoreIpAddresses(s string) []string {
	var res []string
	var bt func(i int, left int, cur []byte)
	bt = func(i int, left int, cur []byte) {
		if left == 0 {
			if valid([]byte(s[i:])) {
				cur = append(cur, s[i:]...)
				res = append(res, string(cur))
			}
			return
		}
		for j := i; j < len(s); j++ {
			if valid([]byte(s[i : j+1])) {
				newCur := append(cur, ([]byte(s[i : j+1]))...)
				bt(j+1, left-1, append(newCur, '.'))
			}
		}
	}

	bt(0, 3, nil)

	return res
}

func valid(s []byte) bool {
	if len(s) > 3 || len(s) == 0 {
		return false
	}
	if s[0] == '0' && len(s) > 1 {
		return false
	}
	i, _ := strconv.ParseInt(string(s), 10, 64)
	return i <= 255
}
