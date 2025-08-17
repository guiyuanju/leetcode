package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(customSortString("cba", "abcd"))
	fmt.Println(customSortString("bcafg", "abcd"))
}

func customSortString(order string, s string) string {
	m := map[byte]int{}
	for i, v := range order {
		m[byte(v)] = i
	}

	res := []byte(s)
	slices.SortFunc(res, func(a, b byte) int {
		o1, ok1 := m[a]
		if !ok1 {
			o1 = int(a)
		}
		o2, ok2 := m[b]
		if !ok2 {
			o2 = int(b)
		}
		return o1 - o2
	})

	return string(res)
}
