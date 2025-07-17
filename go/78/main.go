package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(subsets(nums))

	nums = []int{0}
	fmt.Println(subsets(nums))
}

func subsets(nums []int) [][]int {
	var res [][]int
	var bt func(cur []int, i int)
	bt = func(cur []int, i int) {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		res = append(res, tmp)
		for j := i; j < len(nums); j++ {
			cur = append(cur, nums[j])
			bt(cur, j+1)
			cur = cur[:len(cur)-1]
		}
	}
	bt(nil, 0)
	return res
}
