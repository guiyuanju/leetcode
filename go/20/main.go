package main

import (
	"github.com/guiyuanju/lcutils/assert"
)

func main() {
	assert.Eq(true, isValid("()"))
	assert.Eq(true, isValid("()[]{}"))
	assert.Eq(false, isValid("(]"))
	assert.Eq(true, isValid("([])"))
	assert.Eq(false, isValid("([)]"))
}

func isValid(s string) bool {
	stack := []byte{}
	for _, c := range []byte(s) {
		if c == '(' || c == '[' || c == '{' {
			stack = append(stack, c)
		} else {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			if top == '(' && c == ')' || top == '[' && c == ']' || top == '{' && c == '}' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}
