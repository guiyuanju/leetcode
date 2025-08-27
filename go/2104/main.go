package main

import "fmt"

func main() {
	fmt.Println(subArrayRanges([]int{1, 2, 3}))
	fmt.Println(subArrayRanges([]int{1, 3, 3}))
	fmt.Println(subArrayRanges([]int{4, -2, -3, 4, 1}))
}

func subArrayRanges(nums []int) int64 {
	monoMin := []int{}
	monoMax := []int{}
	mins := make([]int, len(nums))
	maxs := make([]int, len(nums))

	for i, n := range nums {
		for len(monoMin) > 0 && n < nums[monoMin[len(monoMin)-1]] {
			monoMin = monoMin[:len(monoMin)-1]
		}
		monoMin = append(monoMin, i)
		if len(monoMin) == 1 {
			mins[i] = n * (i + 1)
		} else {
			prev := monoMin[len(monoMin)-2]
			mins[i] = mins[prev] + n*(i-prev)
		}

		for len(monoMax) > 0 && n > nums[monoMax[len(monoMax)-1]] {
			monoMax = monoMax[:len(monoMax)-1]
		}
		monoMax = append(monoMax, i)
		if len(monoMax) == 1 {
			maxs[i] = n * (i + 1)
		} else {
			prev := monoMax[len(monoMax)-2]
			maxs[i] = maxs[prev] + n*(i-prev)
		}
	}

	var resMin, resMax int64
	for i := range mins {
		resMin += int64(mins[i])
		resMax += int64(maxs[i])
	}
	return resMax - resMin
}
