package main

import "fmt"

func main() {
	fmt.Println(numSubarraysWithSum([]int{1, 0, 1, 0, 1}, 2))
	fmt.Println(numSubarraysWithSum([]int{0, 0, 0, 0, 0}, 0))
}

func numSubarraysWithSum(nums []int, goal int) int {
	count := map[int]int{}
	count[0] = 1
	var cur, res int
	for _, n := range nums {
		cur += n
		// cur - x = goal; x = cur - goal
		res += count[cur-goal]
		count[cur]++
	}
	return res
}
