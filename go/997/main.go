package main

import "fmt"

func main() {
	trust := [][]int{{1, 2}}
	n := 2
	assertEq(2, findJudge(n, trust))
}

func assertEq[T comparable](a, b T) {
	if a != b {
		fmt.Printf("Failed: %v != %v\n", a, b)
	} else {
		fmt.Printf("Ok: %v\n", a)
	}
}

func findJudge(n int, trust [][]int) int {
	m := map[int][]int{}
	for i := range n {
		m[i] = []int{0, 0}
	}
	for _, t := range trust {
		m[t[0]-1][0]++
		m[t[1]-1][1]++
	}

	for k, v := range m {
		if v[0] == 0 && v[1] == n-1 {
			return k + 1
		}
	}

	return -1
}
