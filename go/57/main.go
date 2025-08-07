package main

import "fmt"

func main() {
	fmt.Println(insert([][]int{{1, 3}, {6, 9}}, []int{2, 5}))
	fmt.Println(insert([][]int{{1, 3}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}))
}

func insert(intervals [][]int, newInterval []int) [][]int {
	var i int
	res := make([][]int, 0, len(intervals)+1)
	for ; i < len(intervals); i++ {
		if newInterval[0] < intervals[i][0] {
			break
		}
		res = append(res, intervals[i])
	}

	if i > 0 && res[len(res)-1][1] >= newInterval[0] {
		res[len(res)-1][1] = max(res[len(res)-1][1], newInterval[1])
	} else {
		res = append(res, newInterval)
	}

	for ; i < len(intervals); i++ {
		last := res[len(res)-1]
		if last[1] >= intervals[i][0] {
			last[1] = max(last[1], intervals[i][1])
		} else {
			res = append(res, intervals[i])
		}
	}

	return res
}
