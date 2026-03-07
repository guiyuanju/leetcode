package main

import (
	"fmt"
	"math"
	"reflect"
	"slices"
)

func main() {
	spells := []int{5, 1, 3}
	potions := []int{1, 2, 3, 4, 5}
	var success int64 = 7
	assertEq([]int{4, 0, 3}, successfulPairs(spells, potions, int64(success)))

	spells = []int{3, 1, 2}
	potions = []int{8, 5, 8}
	success = 16
	assertEq([]int{2, 0, 2}, successfulPairs(spells, potions, int64(success)))
}

func assertEq(a, b any) {
	if reflect.DeepEqual(a, b) {
		fmt.Printf("Ok: %v\n", a)
	} else {
		fmt.Printf("Failed: %v != %v\n", a, b)
	}
}

func successfulPairs(spells []int, potions []int, success int64) []int {
	slices.Sort(potions)
	res := make([]int, len(spells))
	for i := range res {
		res[i] = bs(potions, int(math.Ceil(float64(success)/float64(spells[i]))))
	}
	return res
}

func bs(xs []int, target int) int {
	i := 0
	j := len(xs)
	for i < j {
		mid := i + (j-i)/2
		if xs[mid] < target {
			i = mid + 1
		} else {
			j = mid
		}
	}
	return len(xs) - i
}
