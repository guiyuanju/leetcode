package main

import (
	"fmt"
	"reflect"
	"slices"
)

func main() {
	assertEq(true, makesquare([]int{1, 1, 2, 2, 2}))
	assertEq(false, makesquare([]int{3, 3, 3, 3, 4}))
	assertEq(true, makesquare([]int{5, 5, 5, 5, 4, 4, 4, 4, 3, 3, 3, 3}))
}

func assertEq(a, b any) {
	if reflect.DeepEqual(a, b) {
		fmt.Printf("Ok: %v\n", a)
	} else {
		fmt.Printf("Failed: %v != %v\n", a, b)
	}
}

// Time complexity = O(n * L ^ 4) where L = the max value of a conf[i] = totalSum / 4
// !!! There is a better solution, which is only N * 2 ^ N
func makesquare(matchsticks []int) bool {
	var lenSum int
	for _, s := range matchsticks {
		lenSum += s
	}

	if lenSum%4 != 0 {
		return false
	}

	edgeLen := lenSum / 4

	memo := map[[5]int]bool{}

	var dp func(i int, conf [4]int) bool
	dp = func(i int, conf [4]int) bool {
		slices.Sort(conf[:])

		key := [5]int{i, conf[0], conf[1], conf[2], conf[3]}

		if v, ok := memo[key]; ok {
			return v
		}

		if i >= len(matchsticks) {
			return true
		}

		for j := range conf {
			if conf[j]+matchsticks[i] <= edgeLen {
				conf[j] += matchsticks[i]
				if dp(i+1, conf) {
					memo[key] = true
					return true
				}
				conf[j] -= matchsticks[i]
			}
		}

		memo[key] = false
		return false
	}

	return dp(0, [4]int{0, 0, 0, 0})
}

// func makesquare(matchsticks []int) bool {
// 	var sum int
// 	for _, l := range matchsticks {
// 		sum += l
// 	}
// 	if sum%4 != 0 {
// 		return false
// 	}
// 	width := sum / 4
//
// 	type state struct {
// 		i    int
// 		conf [4]int
// 	}
// 	memo := map[state]bool{}
//
// 	var dp func(i int, conf [4]int) bool
// 	dp = func(i int, conf [4]int) bool {
// 		slices.Sort(conf[:])
// 		if v, ok := memo[state{i, conf}]; ok {
// 			return v
// 		}
//
// 		if i >= len(matchsticks) {
// 			return true
// 		}
//
// 		for j, c := range conf {
// 			if c >= matchsticks[i] {
// 				conf[j] -= matchsticks[i]
// 				if dp(i+1, conf) {
// 					memo[state{i, conf}] = true
// 					return true
// 				}
// 				conf[j] += matchsticks[i]
// 			}
// 		}
//
// 		memo[state{i, conf}] = false
// 		return false
// 	}
//
// 	return dp(0, [4]int{width, width, width, width})
// }

// func makesquare(matchsticks []int) bool {
// 	var sum int
// 	for _, n := range matchsticks {
// 		sum += n
// 	}
// 	if sum%4 != 0 {
// 		return false
// 	}
// 	length := sum / 4
//
// 	type key struct {
// 		i    int
// 		conf [4]int
// 	}
//
// 	memo := map[key]bool{}
//
// 	var dp func(i int, conf [4]int) bool
// 	dp = func(i int, conf [4]int) bool {
// 		slices.Sort(conf[:])
// 		if v, ok := memo[key{i, conf}]; ok {
// 			return v
// 		}
//
// 		if i == len(matchsticks) {
// 			return true
// 		}
//
// 		for j := range conf {
// 			if conf[j]+matchsticks[i] <= length {
// 				conf[j] += matchsticks[i]
// 				if dp(i+1, conf) {
// 					memo[key{i, conf}] = true
// 					return true
// 				}
// 				conf[j] -= matchsticks[i]
// 			}
// 		}
//
// 		memo[key{i, conf}] = false
// 		return false
// 	}
//
// 	return dp(0, [4]int{})
// }
