package main

import "fmt"

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)

	nums = []int{0}
	moveZeroes(nums)
	fmt.Println(nums)
}

func moveZeroes(nums []int) {
	var i int
	for ; i < len(nums); i++ {
		if nums[i] == 0 {
			break
		}
	}

	for j := i + 1; j < len(nums); j++ {
		if nums[j] == 0 {
			continue
		}
		nums[i], nums[j] = nums[j], nums[i]
		i++
	}
}
