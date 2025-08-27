package main

import "fmt"

func main() {
	fmt.Println(mostCompetitive([]int{3, 5, 2, 6}, 2))
	fmt.Println(mostCompetitive([]int{2, 4, 3, 3, 5, 4, 9, 6}, 4))
	fmt.Println(mostCompetitive([]int{3, 4, 5, 1, 0}, 3))
}

func mostCompetitive(nums []int, k int) []int {
	mono := []int{}
	for i, n := range nums {
		for len(mono) > 0 && n < mono[len(mono)-1] && len(mono)+len(nums)-i-1 >= k {
			mono = mono[:len(mono)-1]
		}
		mono = append(mono, n)
	}
	return mono[:k]
}

// [3 4 5 1 0] 3 1 0
