package main

import (
	"fmt"
	"slices"
)

func main() {
	clips := [][]int{{0, 2}, {4, 6}, {8, 10}, {1, 9}, {1, 5}, {5, 9}}
	time := 10
	assertEq(3, videoStitchingSort(clips, time))

	clips = [][]int{{0, 1}, {1, 2}}
	time = 5
	assertEq(-1, videoStitchingSort(clips, time))

	clips = [][]int{{0, 1}, {6, 8}, {0, 2}, {5, 6}, {0, 4}, {0, 3}, {6, 7}, {1, 3}, {4, 7}, {1, 4}, {2, 5}, {2, 6}, {3, 4}, {4, 5}, {5, 7}, {6, 9}}
	time = 9
	assertEq(3, videoStitchingSort(clips, time))

	clips = [][]int{{0, 4}, {2, 8}}
	time = 5
	assertEq(2, videoStitchingSort(clips, time))

	clips = [][]int{{0, 5}, {6, 8}}
	time = 7
	assertEq(-1, videoStitchingSort(clips, time))

	clips = [][]int{{0, 0}, {9, 9}, {2, 10}, {0, 3}, {0, 5}, {3, 4}, {6, 10}, {1, 2}, {4, 7}, {5, 6}}
	time = 5
	assertEq(1, videoStitchingSort(clips, time))
}

func assertEq[T comparable](a, b T) {
	if a != b {
		fmt.Printf("Failed: %v != %v\n", a, b)
	} else {
		fmt.Printf("Ok: %v\n", a)
	}
}

func videoStitchingSort(clips [][]int, time int) int {
	slices.SortFunc(clips, func(a, b []int) int {
		return a[0] - b[0]
	})

	var end, i, res int
	for end < time {
		hi := end
		for i < len(clips) && clips[i][0] <= end {
			hi = max(hi, clips[i][1])
			i++
		}
		if hi == end {
			return -1
		}
		end = hi
		res++
	}
	return res
}

// 最优化问题
// - greedy
// - dynamic programming
// - bianry search
