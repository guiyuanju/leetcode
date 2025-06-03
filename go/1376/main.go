package main

import "fmt"

func main() {
	n := 1
	headId := 0
	manager := []int{-1}
	informTime := []int{0}
	assertEq(0, numOfMinutes(n, headId, manager, informTime))

	n = 6
	headId = 2
	manager = []int{2, 2, -1, 2, 2, 2}
	informTime = []int{0, 0, 1, 0, 0, 0}
	assertEq(1, numOfMinutes(n, headId, manager, informTime))
}

func assertEq[T comparable](a, b T) {
	if a != b {
		fmt.Printf("Failed: %v != %v\n", a, b)
	} else {
		fmt.Printf("Ok: %v\n", a)
	}
}

func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
	totalTime := make([]int, n)
	for i := range n {
		totalTime[i] = -1
	}

	var res int
	for i := range n {
		cur := i
		curTime := 0
		for manager[cur] != -1 {
			if totalTime[cur] != -1 {
				curTime += totalTime[cur]
				break
			}
			curTime += informTime[manager[cur]]
			cur = manager[cur]
		}
		totalTime[i] = curTime
		res = max(res, curTime)
	}

	return res
}
