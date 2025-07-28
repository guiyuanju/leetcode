package main

import "fmt"

func main() {
	fmt.Println(rob([]int{2, 3, 2}))
	fmt.Println(rob([]int{1, 2, 3, 1}))
	fmt.Println(rob([]int{1, 2, 3}))
	fmt.Println(rob([]int{1}))
}

func rob2(nums []int) int {
	var dp func(i int, rob bool) int
	dp = func(i int, rob bool) int {
		if i < 0 {
			return 0
		}
		if i == 0 {
			if rob {
				return 0
			}
			return nums[i]
		}

		return max(nums[i]+dp(i-2, rob), dp(i-1, rob))
	}

	n := len(nums)
	return max(nums[n-1]+dp(n-3, true), dp(n-2, false))
}

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	dp := make([]int, len(nums)+1)
	dp[1] = nums[0]
	for i := 2; i < len(nums); i++ {
		dp[i] = max(nums[i-1]+dp[i-2], dp[i-1])
	}
	res := dp[len(nums)-1]

	clear(dp)
	for i := 2; i < len(nums)+1; i++ {
		dp[i] = max(nums[i-1]+dp[i-2], dp[i-1])
	}
	res = max(res, dp[len(nums)])
	return res
}
