package main

import "fmt"

func main() {
	fmt.Println(sumSubarrayMins([]int{3, 1, 2, 4}))
	fmt.Println(sumSubarrayMins([]int{11, 81, 94, 43, 3}))
}

func sumSubarrayMins(arr []int) int {
	mono := []int{}
	sums := make([]int, len(arr))
	for i, n := range arr {
		for len(mono) > 0 && n < arr[mono[len(mono)-1]] {
			mono = mono[:len(mono)-1]
		}
		mono = append(mono, i)
		if len(mono) == 1 {
			sums[i] = n * (i + 1)
		} else {
			prev := mono[len(mono)-2]
			sums[i] = sums[prev] + n*(i-prev)
		}
	}

	var res int
	for _, n := range sums {
		res += n
		res %= 1e9 + 7
	}
	return res
}
