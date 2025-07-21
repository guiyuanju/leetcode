package main

import (
	"github.com/guiyuanju/lcutils/assert"
)

func main() {
	assert.Eq("c", getHappyString(1, 3))
	assert.Eq("", getHappyString(1, 4))
	assert.Eq("cab", getHappyString(3, 9))
}

func getHappyString(n int, k int) string {
	var res string
	var count int
	set := []byte{'a', 'b', 'c'}
	var bt func(cur []byte) bool
	bt = func(cur []byte) bool {
		if len(cur) == n {
			count++
			if count == k {
				res = string(cur)
				return true
			}
			return false
		}

		for _, b := range set {
			if len(cur) > 0 && cur[len(cur)-1] == b {
				continue
			}
			if bt(append(cur, b)) {
				return true
			}
		}

		return false
	}

	bt(nil)

	return res
}
