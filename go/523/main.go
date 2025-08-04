package main

import "fmt"

func main() {
	fmt.Println(checkSubarraySum([]int{23, 2, 4, 6, 7}, 6))
	fmt.Println(checkSubarraySum([]int{23, 2, 6, 4, 7}, 6))
	fmt.Println(checkSubarraySum([]int{23, 2, 6, 4, 7}, 13))
	fmt.Println(checkSubarraySum([]int{1, 0}, 2))
	fmt.Println(checkSubarraySum([]int{5, 0, 0, 0}, 3))
}

func checkSubarraySum(nums []int, k int) bool {
	if len(nums) < 2 {
		return false
	}

	m := map[int]int{}
	m[0] = -1

	var cur int
	for i, n := range nums {
		cur += n
		if idx, ok := m[cur%k]; ok {
			if i-idx > 1 {
				return true
			}
		} else {
			m[cur%k] = i
		}
	}

	return false
}
