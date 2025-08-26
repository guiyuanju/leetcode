package main

import "fmt"

func main() {
	fmt.Println(finalPrices([]int{8, 4, 6, 2, 3}))
	fmt.Println(finalPrices([]int{1, 2, 3, 4, 5}))
	fmt.Println(finalPrices([]int{10, 1, 1, 6}))
}

func finalPrices(prices []int) []int {
	res := make([]int, len(prices))
	copy(res, prices)
	mono := []int{}
	for i, p := range prices {
		for len(mono) > 0 && p <= prices[mono[len(mono)-1]] {
			res[mono[len(mono)-1]] = prices[mono[len(mono)-1]] - p
			mono = mono[:len(mono)-1]
		}
		mono = append(mono, i)
	}
	return res
}
