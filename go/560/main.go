package main

import "fmt"

func main() {
	fmt.Println(subarraySum([]int{1, 1, 1}, 2))
	fmt.Println(subarraySum([]int{1, 2, 3}, 3))
}

func subarraySum(nums []int, k int) int {
	count := map[int]int{}
	count[0] = 1
	var cur, res int
	for i := range nums {
		cur += nums[i]
		res += count[cur-k]
		count[cur]++
	}
	return res
}
