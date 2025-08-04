package main

import (
	"fmt"

	"github.com/guiyuanju/lcutils/assert"
)

func main() {
	assert.Eq("Bob", stoneGameIII([]int{1, 2, 3, 7}))
	assert.Eq("Alice", stoneGameIII([]int{1, 2, 3, -9}))
	assert.Eq("Tie", stoneGameIII([]int{1, 2, 3, 6}))
	assert.Eq("Tie", stoneGameIII([]int{-1, -2, -3}))
}

func stoneGameIII(stoneValue []int) string {
	type state struct {
		alice bool
		i     int
	}
	memo := map[state][]int{}

	var dp func(alice bool, i int) (aliceScore int, bobScore int)
	dp = func(alice bool, i int) (int, int) {
		if v, ok := memo[state{alice, i}]; ok {
			return v[0], v[1]
		}

		if i >= len(stoneValue) {
			return 0, 0
		}

		as1, bs1 := dp(!alice, i+1)
		as2, bs2 := dp(!alice, i+2)
		as3, bs3 := dp(!alice, i+3)

		results := [][]int{{as1, bs1}, {as2, bs2}, {as3, bs3}}
		if alice {
			var sum int
			for j := 0; j < 3 && i+j < len(stoneValue); j++ {
				sum += stoneValue[i+j]
				results[j][0] += sum
			}
			maximum := results[0]
			for j := 0; j < 3 && i+j < len(stoneValue); j++ {
				if results[j][0] > maximum[0] {
					maximum = results[j]
				}
			}
			memo[state{alice, i}] = maximum
			return maximum[0], maximum[1]
		} else {
			var sum int
			for j := 0; j < 3 && i+j < len(stoneValue); j++ {
				sum += stoneValue[i+j]
				results[j][1] += sum
			}
			maximum := results[0]
			for j := 0; j < 3 && i+j < len(stoneValue); j++ {
				if results[j][1] > maximum[1] {
					maximum = results[j]
				}
			}
			memo[state{alice, i}] = maximum
			return maximum[0], maximum[1]
		}
	}

	a, b := dp(true, 0)
	fmt.Println(a, b)
	if a > b {
		return "Alice"
	} else if a < b {
		return "Bob"
	} else {
		return "Tie"
	}
}
