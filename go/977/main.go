package main

import "fmt"

func main() {
	nums := []int{-4, -1, 0, 3, 10}
	fmt.Println(sortedSquares(nums))
}

func sortedSquares(nums []int) []int {
	res := make([]int, len(nums))
	i := 0
	j := len(nums) - 1
	cur := j
	for i <= j {
		if abs(nums[i]) >= abs(nums[j]) {
			res[cur] = nums[i] * nums[i]
			i++
		} else {
			res[cur] = nums[j] * nums[j]
			j--
		}
		cur--
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
