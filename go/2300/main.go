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
	var res []int
	slices.Sort(potions)
	for _, s := range spells {
		target := int64(math.Ceil(float64(success) / float64(s)))
		res = append(res, len(potions)-bs(potions, target))
	}
	return res
}

func bs(xs []int, target int64) int {
	lo := 0
	hi := len(xs)
	for lo < hi {
		mid := lo + (hi-lo)/2
		if int64(xs[mid]) < target {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}
