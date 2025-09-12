package main

import (
	"fmt"
	"slices"
)

func main() {
	spells := []int{5, 1, 3}
	potions := []int{1, 2, 3, 4, 5}
	var success int64 = 7
	fmt.Println(successfulPairs(spells, potions, int64(success)))

	spells = []int{3, 1, 2}
	potions = []int{8, 5, 8}
	success = 16
	fmt.Println(successfulPairs(spells, potions, int64(success)))
}

func successfulPairs(spells []int, potions []int, success int64) []int {
	slices.Sort(potions)
	res := make([]int, len(spells))
	for i := range res {
		res[i] = len(potions) - bs(potions, float64(success)/float64(spells[i]))
	}
	return res
}

func bs(arr []int, target float64) int {
	i := 0
	j := len(arr)
	for i < j {
		mid := i + (j-i)/2
		if float64(arr[mid]) >= target {
			j = mid
		} else {
			i = mid + 1
		}
	}
	return i
}
