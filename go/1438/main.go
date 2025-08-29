package main

import "fmt"

func main() {
	fmt.Println(longestSubarray([]int{8, 2, 4, 7}, 4))
	fmt.Println(longestSubarray([]int{10, 1, 2, 4, 7, 2}, 5))
	fmt.Println(longestSubarray([]int{4, 2, 2, 2, 4, 4, 2, 2}, 0))
}

func longestSubarray(nums []int, limit int) int {
	monoMin := []int{}
	monoMax := []int{}
	var res int
	for i, j := 0, 0; j < len(nums); j++ {
		for len(monoMin) > 0 && nums[j] < monoMin[len(monoMin)-1] {
			monoMin = monoMin[:len(monoMin)-1]
		}
		monoMin = append(monoMin, nums[j])
		for len(monoMax) > 0 && nums[j] > monoMax[len(monoMax)-1] {
			monoMax = monoMax[:len(monoMax)-1]
		}
		monoMax = append(monoMax, nums[j])
		for i <= j && monoMax[0]-monoMin[0] > limit {
			if monoMin[0] == nums[i] {
				monoMin = monoMin[1:]
			}
			if monoMax[0] == nums[i] {
				monoMax = monoMax[1:]
			}
			i++
		}
		if len(monoMin) > 0 && len(monoMax) > 0 {
			res = max(res, j-i+1)
		}
	}
	return res
}
