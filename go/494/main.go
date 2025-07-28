package main

import "fmt"

func main() {
	fmt.Println(findTargetSumWays([]int{1, 1, 1, 1, 1}, 3))
	fmt.Println(findTargetSumWays([]int{1}, 1))
}

func findTargetSumWays(nums []int, target int) int {
	memo := map[[2]int]int{}
	var dp func(i int, expect int) int
	dp = func(i int, expect int) int {
		if v, ok := memo[[2]int{i, expect}]; ok {
			return v
		}

		if i == len(nums) {
			if expect == 0 {
				return 1
			}
			return 0
		}

		memo[[2]int{i, expect}] = dp(i+1, expect-nums[i]) + dp(i+1, expect+nums[i])
		return memo[[2]int{i, expect}]
	}

	return dp(0, target)
}
