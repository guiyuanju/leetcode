package main

import "fmt"

func main() {
	fmt.Println(getSmallestString(3, 27))
	fmt.Println(getSmallestString(5, 73))
}

func getSmallestString(n int, k int) string {
	res := make([]byte, n)
	for i := range n {
		res[i] = 'a'
	}
	left := k - n
	for i := n - 1; i >= 0; i-- {
		if k == 0 {
			break
		}
		if left < 26 {
			res[i] = byte('a' + left)
			break
		}
		res[i] = 'z'
		left -= 'z' - 'a'
	}
	return string(res)
}
