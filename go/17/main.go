package main

import (
	"strconv"

	"github.com/guiyuanju/lcutils/assert"
)

func main() {
	digits := "23"
	assert.Eq([]string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}, letterCombinations(digits))

	digits = ""
	assert.Eq(0, len(letterCombinations(digits)))
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}

	m := []string{
		"abc",
		"def",
		"ghi",
		"jkl",
		"mno",
		"pqrs",
		"tuv",
		"wxyz",
	}

	var res []string
	var bt func(cur string, i int)
	bt = func(cur string, i int) {
		if i == len(digits) {
			res = append(res, cur)
			return
		}
		index, _ := strconv.Atoi(string(digits[i]))
		for _, b := range m[index-2] {
			cur += string(b)
			bt(cur, i+1)
			cur = cur[:len(cur)-1]
		}
	}
	bt("", 0)
	return res
}
