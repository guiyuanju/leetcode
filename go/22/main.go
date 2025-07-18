package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(3))
	fmt.Println(generateParenthesis(1))
}

func generateParenthesis(n int) []string {
	var res []string
	var bt func(cur string, left, right int)
	bt = func(cur string, left, right int) {
		if len(cur) == n*2 {
			res = append(res, cur)
			return
		}
		if left < n {
			bt(cur+"(", left+1, right)
		}
		if left > right {
			bt(cur+")", left, right+1)
		}
	}

	bt("", 0, 0)
	return res
}
