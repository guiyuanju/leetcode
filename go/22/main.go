package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(3))
	fmt.Println(generateParenthesis(1))
}

func generateParenthesis(n int) []string {
	var res []string
	var bt func(cur []byte, left int, right int)
	bt = func(cur []byte, left int, right int) {
		if left == n && right == n {
			tmp := make([]byte, len(cur))
			copy(tmp, cur)
			res = append(res, string(tmp))
			return
		}

		if left < n {
			bt(append(cur, '('), left+1, right)
		}
		if right < left {
			bt(append(cur, ')'), left, right+1)
		}
	}

	bt(nil, 0, 0)

	return res
}
