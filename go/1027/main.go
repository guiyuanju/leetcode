package main

import "fmt"

func main() {
	fmt.Println(longestArithSeqLength([]int{3, 6, 9, 12}))
	fmt.Println(longestArithSeqLength([]int{9, 4, 7, 2, 10}))
	fmt.Println(longestArithSeqLength([]int{20, 1, 15, 3, 10, 5, 8}))
}

func longestArithSeqLength(nums []int) int {
	idx := make(map[int][]int, len(nums))
	for i, n := range nums {
		idx[n] = append(idx[n], i)
	}

	memo := map[[2]int]int{}

	var dp func(i, diff int) int
	dp = func(i, diff int) int {
		if v, ok := memo[[2]int{i, diff}]; ok {
			return v
		}

		res := 1
		for _, j := range idx[nums[i]-diff] {
			if j < i {
				res = max(res, 1+dp(j, diff))
			}
		}

		memo[[2]int{i, diff}] = res
		return res
	}

	var res int
	for i := range nums {
		for j := range i {
			res = max(res, 1+dp(j, nums[i]-nums[j]))
		}
	}
	return res
}
