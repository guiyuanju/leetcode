package main

import (
	"math"

	"github.com/guiyuanju/lcutils/assert"
)

func main() {
	assert.Eq("Bob", stoneGameIII([]int{1, 2, 3, 7}))
	assert.Eq("Alice", stoneGameIII([]int{1, 2, 3, -9}))
	assert.Eq("Tie", stoneGameIII([]int{1, 2, 3, 6}))
	assert.Eq("Tie", stoneGameIII([]int{-1, -2, -3}))
	assert.Eq("Bob", stoneGameIII([]int{-2}))
}

func stoneGameIII(stoneValue []int) string {
	memo := map[[2]int][2]int{}
	var dp func(i int, who int) [2]int // 0 - alice, 1 - bob
	dp = func(i int, who int) [2]int {
		if v, ok := memo[[2]int{i, who}]; ok {
			return v
		}

		if i >= len(stoneValue) {
			return [2]int{0, 0}
		}

		var toTake int
		res := [2]int{math.MinInt, math.MinInt}
		for j := range 3 {
			if i+j < len(stoneValue) {
				cur := dp(i+j+1, 1-who)
				toTake += stoneValue[i+j]
				if cur[who]+toTake > res[who] {
					res[who] = cur[who] + toTake
					res[1-who] = cur[1-who]
				}
			}
		}

		memo[[2]int{i, who}] = res

		return res
	}

	res := dp(0, 0)
	if res[0] == res[1] {
		return "Tie"
	} else if res[0] > res[1] {
		return "Alice"
	} else {
		return "Bob"
	}
}
