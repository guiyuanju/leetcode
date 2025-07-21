package main

import "fmt"

func main() {
	nums := []int{4, 6, 7, 7}
	fmt.Println(findSubsequences(nums))

	nums = []int{4, 4, 3, 2, 1}
	fmt.Println(findSubsequences(nums))
}

func findSubsequences(nums []int) [][]int {
	var res [][]int
	var bt func(cur []int, i int)
	bt = func(cur []int, i int) {
		if len(cur) > 1 {
			res = append(res, append([]int(nil), cur...))
		}

		selected := make(map[int]bool, len(nums)-i)
		for j := i; j < len(nums); j++ {
			if len(cur) > 0 && nums[j] < cur[len(cur)-1] {
				continue
			}
			if selected[nums[j]] {
				continue
			}
			selected[nums[j]] = true
			bt(append(cur, nums[j]), j+1)
		}
	}

	bt(nil, 0)

	return res
}
