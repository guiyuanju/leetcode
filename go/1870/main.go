package main

import (
	"fmt"
	"math"
)

func main() {
	dist := []int{1, 3, 2}
	assertEq(1, minSpeedOnTime(dist, 6))

	dist = []int{1, 3, 2}
	assertEq(3, minSpeedOnTime(dist, 2.7))

	dist = []int{1, 3, 2}
	assertEq(-1, minSpeedOnTime(dist, 1.9))

	dist = []int{1, 1, 100000}
	assertEq(10000000, minSpeedOnTime(dist, 2.01))
}

func assertEq[T comparable](a, b T) {
	if a != b {
		fmt.Printf("Failed: %v != %v\n", a, b)
	} else {
		fmt.Printf("Ok: %v\n", a)
	}
}

func minSpeedOnTime(dist []int, hour float64) int {
	check := func(speed int) bool {
		var res float64
		for _, d := range dist[:len(dist)-1] {
			res += math.Ceil(float64(d) / float64(speed))
		}
		res += float64(dist[len(dist)-1]) / float64(speed)
		return res <= hour
	}

	var left, right int
	right = 1e7

	for left <= right {
		mid := left + (right-left)/2
		if check(mid) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if left > 1e7 {
		return -1
	}
	return left
}
