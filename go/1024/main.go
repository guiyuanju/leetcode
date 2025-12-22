package main

import (
	"fmt"
	"slices"
)

func main() {
	clips := [][]int{{0, 2}, {4, 6}, {8, 10}, {1, 9}, {1, 5}, {5, 9}}
	time := 10
	assertEq(3, videoStitching(clips, time))

	clips = [][]int{{0, 1}, {1, 2}}
	time = 5
	assertEq(-1, videoStitching(clips, time))

	clips = [][]int{{0, 1}, {6, 8}, {0, 2}, {5, 6}, {0, 4}, {0, 3}, {6, 7}, {1, 3}, {4, 7}, {1, 4}, {2, 5}, {2, 6}, {3, 4}, {4, 5}, {5, 7}, {6, 9}}
	time = 9
	assertEq(3, videoStitching(clips, time))

	clips = [][]int{{0, 4}, {2, 8}}
	time = 5
	assertEq(2, videoStitching(clips, time))

	clips = [][]int{{0, 5}, {6, 8}}
	time = 7
	assertEq(-1, videoStitching(clips, time))

	clips = [][]int{{0, 0}, {9, 9}, {2, 10}, {0, 3}, {0, 5}, {3, 4}, {6, 10}, {1, 2}, {4, 7}, {5, 6}}
	time = 5
	assertEq(1, videoStitching(clips, time))
}

func assertEq[T comparable](a, b T) {
	if a != b {
		fmt.Printf("Failed: %v != %v\n", a, b)
	} else {
		fmt.Printf("Ok: %v\n", a)
	}
}

func videoStitching(clips [][]int, time int) int {
	var res, curLimit, nextLimit int
	res = 1
	slices.SortFunc(clips, func(a, b []int) int { return a[0] - b[0] })
	for _, c := range clips {
		if nextLimit >= time {
			return res
		}
		if c[0] > nextLimit {
			return -1
		}
		if c[0] > curLimit && c[1] <= nextLimit {
			continue
		}
		if c[0] > curLimit {
			res++
			curLimit = nextLimit
		}
		nextLimit = max(nextLimit, c[1])
	}
	if nextLimit < time {
		return -1
	}
	return res
}
