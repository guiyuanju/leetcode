package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(makesquare([]int{1, 1, 2, 2, 2}))
	fmt.Println(makesquare([]int{3, 3, 3, 3, 4}))
	fmt.Println(makesquare([]int{5, 5, 5, 5, 4, 4, 4, 4, 3, 3, 3, 3}))
}

// backtracking
// time limit exceeded
func makesquareBt(matchsticks []int) bool {
	var sum int
	for _, n := range matchsticks {
		sum += n
	}
	if sum%4 != 0 {
		return false
	}
	length := sum / 4
	slices.Sort(matchsticks)
	edges := [4]int{}

	var bt func(i int) bool
	bt = func(i int) bool {
		if i == len(matchsticks) {
			return true
		}

		for j := range 4 {
			if edges[j]+matchsticks[i] > length {
				continue
			}
			edges[j] += matchsticks[i]
			if bt(i + 1) {
				return true
			}
			edges[j] -= matchsticks[i]
		}
		return false
	}

	return bt(0)
}

func makesquare(matchsticks []int) bool {
	var sum int
	for _, n := range matchsticks {
		sum += n
	}
	if sum%4 != 0 {
		return false
	}
	length := sum / 4

	type key struct {
		i    int
		conf [4]int
	}

	memo := map[key]bool{}

	var dp func(i int, conf [4]int) bool
	dp = func(i int, conf [4]int) bool {
		slices.Sort(conf[:])
		if v, ok := memo[key{i, conf}]; ok {
			return v
		}

		if i == len(matchsticks) {
			return true
		}

		for j := range conf {
			if conf[j]+matchsticks[i] <= length {
				conf[j] += matchsticks[i]
				if dp(i+1, conf) {
					memo[key{i, conf}] = true
					return true
				}
				conf[j] -= matchsticks[i]
			}
		}

		memo[key{i, conf}] = false
		return false
	}

	return dp(0, [4]int{})
}
