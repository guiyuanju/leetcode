package main

import "fmt"

func main() {
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	fmt.Println(maxSlidingWindow([]int{1}, 1))
}

func maxSlidingWindow(nums []int, k int) []int {
	var res []int
	var i int
	mono := []int{}
	for j, n := range nums {
		for len(mono) > 0 && n > mono[len(mono)-1] {
			mono = mono[:len(mono)-1]
		}
		mono = append(mono, n)
		if j-i+1 == k {
			res = append(res, mono[0])
			if mono[0] == nums[i] {
				mono = mono[1:]
			}
			i++
		}
	}
	return res
}
