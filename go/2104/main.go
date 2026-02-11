package main

import (
	"fmt"
	"reflect"
)

func main() {
	assertEq(int64(4), subArrayRanges([]int{1, 2, 3}))
	assertEq(int64(4), subArrayRanges([]int{1, 3, 3}))
	assertEq(int64(59), subArrayRanges([]int{4, -2, -3, 4, 1}))
}

func assertEq(a, b any) {
	if reflect.DeepEqual(a, b) {
		fmt.Printf("Ok: %v\n", a)
	} else {
		fmt.Printf("Failed: %v != %v\n", a, b)
	}
}

func subArrayRanges(nums []int) int64 {
	mins := make([]int, len(nums))
	maxs := make([]int, len(nums))
	monoMins := []int{}
	monoMaxs := []int{}

	for i, n := range nums {
		for len(monoMins) > 0 && n < nums[monoMins[len(monoMins)-1]] {
			monoMins = monoMins[:len(monoMins)-1]
		}
		monoMins = append(monoMins, i)
		if len(monoMins) == 1 {
			mins[i] = n * (i + 1)
		} else {
			prevIdx := monoMins[len(monoMins)-2]
			mins[i] = mins[prevIdx] + n*(i-prevIdx)
		}

		for len(monoMaxs) > 0 && n > nums[monoMaxs[len(monoMaxs)-1]] {
			monoMaxs = monoMaxs[:len(monoMaxs)-1]
		}
		monoMaxs = append(monoMaxs, i)
		if len(monoMaxs) == 1 {
			maxs[i] = n * (i + 1)
		} else {
			prevIdx := monoMaxs[len(monoMaxs)-2]
			maxs[i] = maxs[prevIdx] + n*(i-prevIdx)
		}
	}

	var res int64
	for i := range mins {
		res += int64(maxs[i]) - int64(mins[i])
	}

	return res
}
