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
	var res []int
	for _, s := range spells {
		target := int(ceil(success, int64(s)))
		idx := bs(potions, target)
		res = append(res, len(potions)-idx)
	}
	return res
}

func ceil(x, y int64) int64 {
	if x%y > 0 {
		return x/y + 1
	}
	return x / y
}

func bs(potions []int, target int) int {
	left := 0
	right := len(potions)
	for left < right {
		mid := left + (right-left)/2
		if potions[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
