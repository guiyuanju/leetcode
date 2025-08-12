package main

import "fmt"

func main() {
	fmt.Println(longestOnes([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2))
	fmt.Println(longestOnes([]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3))
}

func longestOnes(nums []int, k int) int {
	var i, j, cur, res int
	for ; j < len(nums); j++ {
		if nums[j] == 0 {
			cur++
		}
		for cur > k && i <= j {
			if nums[i] == 0 {
				cur--
			}
			i++
		}
		res = max(res, j-i+1)
	}
	return res
}
