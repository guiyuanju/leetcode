package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(rob(nums))

	nums = []int{2, 7, 9, 3, 1}
	fmt.Println(rob(nums))
}

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	dp0 := nums[0]
	dp1 := max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		dp0 = max(nums[i]+dp0, dp1)
		dp0, dp1 = dp1, dp0
	}

	return dp1
}
