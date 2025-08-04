package main

import (
	"math"
	"slices"

	"github.com/guiyuanju/lcutils/assert"
)

func main() {
	assert.Eq(false, carPooling3([][]int{{2, 1, 5}, {3, 3, 7}}, 4))
	assert.Eq(true, carPooling3([][]int{{2, 1, 5}, {3, 3, 7}}, 5))
	assert.Eq(true, carPooling3([][]int{{3, 2, 9}, {3, 2, 8}, {5, 2, 5}, {7, 4, 8}, {8, 7, 8}, {9, 1, 2}}, 22))
}

// O(n*log(n)), better than solution 2
func carPooling3(trips [][]int, capacity int) bool {
	diff := make([][]int, 0, len(trips)*2)
	for _, t := range trips {
		diff = append(diff, []int{t[1], t[0]})
		diff = append(diff, []int{t[2], -t[0]})
	}

	slices.SortFunc(diff, func(a, b []int) int { return a[0] - b[0] })

	var cur int
	for _, d := range diff {
		cur += d[1]
		if cur > capacity {
			return false
		}
	}

	return true
}

// O(n*log(n))
func carPooling2(trips [][]int, capacity int) bool {
	starts := slices.SortedFunc(slices.Values(trips), func(a, b []int) int { return a[1] - b[1] })
	ends := slices.SortedFunc(slices.Values(trips), func(a, b []int) int { return a[2] - b[2] })

	var i, j, cur int
	for i < len(trips) && j < len(trips) {
		if starts[i][1] < ends[j][2] {
			cur += starts[i][0]
			i++
		} else if starts[i][1] > ends[j][2] {
			cur -= ends[j][0]
			j++
		} else {
			cur += starts[i][0]
			cur -= ends[j][0]
			i++
			j++
		}

		if cur > capacity {
			return false
		}
	}

	return true
}

// difference array, better time & space, O(m + n)
func carPooling(trips [][]int, capacity int) bool {
	smallest := math.MaxInt
	largest := math.MinInt
	for _, t := range trips {
		smallest = min(smallest, t[2])
		largest = max(largest, t[2])
	}

	diff := make([]int, largest+1)
	for _, t := range trips {
		diff[t[1]] += t[0]
		diff[t[2]] -= t[0]
	}

	var cur int
	for _, d := range diff {
		cur += d
		if cur > capacity {
			return false
		}
	}

	return true
}
