package main

import (
	"fmt"
	"reflect"
)

func main() {
	assertEq(17, sumSubarrayMins([]int{3, 1, 2, 4}))
	assertEq(444, sumSubarrayMins([]int{11, 81, 94, 43, 3}))
}

func assertEq(a, b any) {
	if reflect.DeepEqual(a, b) {
		fmt.Printf("Ok: %v\n", a)
	} else {
		fmt.Printf("Failed: %v != %v\n", a, b)
	}
}

func sumSubarrayMins(arr []int) int {
	var res int
	mono := []int{}
	sums := make([]int, len(arr))
	for i, n := range arr {
		for len(mono) > 0 && n <= arr[mono[len(mono)-1]] {
			mono = mono[:len(mono)-1]
		}
		mono = append(mono, i)

		if len(mono) == 1 {
			sums[i] = n * (i + 1)
		} else {
			prev := mono[len(mono)-2]
			sums[i] = sums[prev] + n*(i-prev)
		}

		res += sums[i]
		res %= 1e9 + 7
	}

	return res
}

// func sumSubarrayMins(arr []int) int {
// 	mono := []int{}
// 	sums := make([]int, len(arr))
// 	for i, n := range arr {
// 		for len(mono) > 0 && n < arr[mono[len(mono)-1]] {
// 			mono = mono[:len(mono)-1]
// 		}
// 		mono = append(mono, i)
// 		if len(mono) == 1 {
// 			sums[i] = n * (i + 1)
// 		} else {
// 			prev := mono[len(mono)-2]
// 			sums[i] = sums[prev] + n*(i-prev)
// 		}
// 	}
//
// 	var res int
// 	for _, n := range sums {
// 		res += n
// 		res %= 1e9 + 7
// 	}
// 	return res
// }
//
