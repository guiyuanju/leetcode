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
	time := func(speed int) float64 {
		var res float64
		for i := range len(dist) - 1 {
			res += math.Ceil(float64(dist[i]) / float64(speed))
		}
		return res + float64(dist[len(dist)-1])/float64(speed)
	}

	lo := 1
	var hi int = 1e7 + 1
	old := hi

	for lo < hi {
		mid := lo + (hi-lo)/2
		// fmt.Println(lo, hi, mid, time(mid))
		if time(mid) <= hour {
			hi = mid
		} else {
			lo = mid + 1
		}
	}

	if hi == old {
		return -1
	}
	return lo
}
