package main

import "fmt"

func main() {
	fmt.Println(maximumUniqueSubarray([]int{4, 2, 4, 5, 6}))
	fmt.Println(maximumUniqueSubarray([]int{5, 2, 1, 2, 5, 2, 1, 2, 5}))
}

func maximumUniqueSubarray(nums []int) int {
	exist := map[int]bool{}
	var i, j, cur, res int
	for ; j < len(nums); j++ {
		cur += nums[j]
		for exist[nums[j]] {
			exist[nums[i]] = false
			cur -= nums[i]
			i++
		}
		exist[nums[j]] = true
		res = max(res, cur)
	}
	return res
}
