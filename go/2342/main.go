package main

import "fmt"

func main() {
	fmt.Println(maximumSum([]int{18, 43, 36, 13, 7}))
	fmt.Println(maximumSum([]int{10, 12, 19, 14}))
}

func maximumSum(nums []int) int {
	m := map[int]int{}
	res := -1
	for _, n := range nums {
		cur := sumOfDigits(n)
		if v, ok := m[cur]; ok {
			res = max(res, v+n)
		}
		m[cur] = max(n, m[cur])
	}
	return res
}

func sumOfDigits(x int) int {
	var res int
	for x > 0 {
		res += x % 10
		x /= 10
	}
	return res
}
