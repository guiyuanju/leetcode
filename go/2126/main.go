package main

import "fmt"
import "slices"

func main() {
	mass := 10
	asteroids := []int{3, 9, 19, 5, 21}
	assertEq(true, asteroidsDestroyed(mass, asteroids))

	mass = 5
	asteroids = []int{4, 9, 23, 4}
	assertEq(false, asteroidsDestroyed(mass, asteroids))
}

func assertEq[T comparable](a, b T) {
	if a != b {
		fmt.Printf("Failed: %v != %v\n", a, b)
	} else {
		fmt.Printf("Ok: %v\n", a)
	}
}

func asteroidsDestroyed(mass int, asteroids []int) bool {
	slices.Sort(asteroids)
	for _, cur := range asteroids {
		if cur <= mass {
			mass += cur
		} else {
			return false
		}
	}
	return true
}
