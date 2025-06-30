package main

import (
	"fmt"
	"slices"
)

func main() {
	clips := [][]int{{0, 2}, {4, 6}, {8, 10}, {1, 9}, {1, 5}, {5, 9}}
	time := 10
	assertEq(3, videoStitchingDp(clips, time))

	clips = [][]int{{0, 1}, {1, 2}}
	time = 5
	assertEq(-1, videoStitchingDp(clips, time))

	clips = [][]int{{0, 1}, {6, 8}, {0, 2}, {5, 6}, {0, 4}, {0, 3}, {6, 7}, {1, 3}, {4, 7}, {1, 4}, {2, 5}, {2, 6}, {3, 4}, {4, 5}, {5, 7}, {6, 9}}
	time = 9
	assertEq(3, videoStitchingDp(clips, time))

	clips = [][]int{{0, 4}, {2, 8}}
	time = 5
	assertEq(2, videoStitchingDp(clips, time))

	clips = [][]int{{0, 5}, {6, 8}}
	time = 7
	assertEq(-1, videoStitchingDp(clips, time))

	clips = [][]int{{0, 0}, {9, 9}, {2, 10}, {0, 3}, {0, 5}, {3, 4}, {6, 10}, {1, 2}, {4, 7}, {5, 6}}
	time = 5
	assertEq(1, videoStitchingDp(clips, time))
}

func assertEq[T comparable](a, b T) {
	if a != b {
		fmt.Printf("Failed: %v != %v\n", a, b)
	} else {
		fmt.Printf("Ok: %v\n", a)
	}
}

func videoStitchingSort(clips [][]int, time int) int {
	slices.SortFunc(clips, func(a, b []int) int { return a[0] - b[0] })
	var i, end, res int
	for end < time {
		tmp := end
		for ; i < len(clips) && clips[i][0] <= end; i++ {
			tmp = max(tmp, clips[i][1])
		}
		if tmp == end {
			return -1
		}
		end = tmp
		res++
	}
	return res
}

func videoStitchingDp(clips [][]int, time int) int {
	dp := slices.Repeat([]int{time + 1}, time+1)
	dp[0] = 0
	for i := 1; i < len(dp); i++ {
		for _, c := range clips {
			if c[0] >= i || c[1] < i {
				continue
			}
			dp[i] = min(dp[i], dp[c[0]]+1)
		}
		if dp[i] >= time+1 {
			return -1
		}
	}
	return dp[len(dp)-1]
}
