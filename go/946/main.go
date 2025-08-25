package main

import "fmt"

func main() {
	fmt.Println(validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1}))
	fmt.Println(validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 3, 5, 1, 2}))
}

func validateStackSequences(pushed []int, popped []int) bool {
	stack := []int{}
	for _, n := range pushed {
		stack = append(stack, n)
		for len(stack) > 0 && len(popped) > 0 && popped[0] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			popped = popped[1:]
		}
	}
	return len(stack) == 0
}
