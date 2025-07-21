package main

import (
	"fmt"
	"strconv"
	"strings"
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
	var bt func(cur []string, i int)
	bt = func(cur []string, i int) {
		if len(cur) == 4 {
			if i == len(s) {
				res = append(res, strings.Join(cur, "."))
			}
			return
		}
		if i >= len(s) {
			return
		}

		if s[i] == '0' {
			bt(append(cur, "0"), i+1)
			return
		}

		for j := i; j < len(s); j++ {
			part := s[i : j+1]
			if !valid(part) {
				break
			}
			bt(append(cur, part), j+1)
		}
	}

	bt(nil, 0)
	return res
}

func valid(s string) bool {
	n, _ := strconv.Atoi(s)
	return n <= 255
}
