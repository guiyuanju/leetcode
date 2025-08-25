package main

import "fmt"

func main() {
	fmt.Println(robotWithString("zza"))
	fmt.Println(robotWithString("bac"))
	fmt.Println(robotWithString("bdda"))
}

func robotWithString(s string) string {
	mins := make([]byte, len(s))
	mins[len(mins)-1] = s[len(mins)-1]
	for i := len(s) - 2; i >= 0; i-- {
		mins[i] = min(mins[i+1], s[i])
	}

	i := 1
	stack := []byte{s[0]}
	var res []byte
	for i < len(s) {
		if len(stack) > 0 && stack[len(stack)-1] <= mins[i] {
			res = append(res, stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
			i++
		}
	}

	for i := len(stack) - 1; i >= 0; i-- {
		res = append(res, stack[i])
	}

	return string(res)
}
