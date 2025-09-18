package main

import (
	"fmt"
	"reflect"
	"slices"
)

func main() {
	price := []int{13, 5, 1, 8, 21, 2}
	assertEq(8, maximumTastiness(price, 3))

	price = []int{1, 3, 1}
	assertEq(2, maximumTastiness(price, 2))

	price = []int{7, 7, 7, 7}
	assertEq(0, maximumTastiness(price, 2))
}

func assertEq(a, b any) {
	if reflect.DeepEqual(a, b) {
		fmt.Printf("Ok: %v\n", a)
	} else {
		fmt.Printf("Failed: %v != %v\n", a, b)
	}
}

func maximumTastiness(price []int, k int) int {
	slices.Sort(price)
	check := func(g int) bool {
		cur := price[0]
		count := 1
		for i := 1; i < len(price); i++ {
			if price[i]-cur >= g {
				cur = price[i]
				count++
			}
		}
		return count >= k
	}

	lo := 0
	hi := price[len(price)-1] - price[0] + 1
	for lo < hi {
		mid := lo + (hi-lo)/2
		if check(mid) {
			lo = mid + 1
		} else {
			hi = mid
		}
	}

	return lo - 1
}
