package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	matrix := parse2DArray("[[1,3,5,7],[10,11,16,20],[23,30,34,60]]")
	target := 3
	assertEq(true, searchMatrix(matrix, target))

	matrix = parse2DArray("[[1,3,5,7],[10,11,16,20],[23,30,34,60]]")
	target = 13
	assertEq(false, searchMatrix(matrix, target))

	matrix = parse2DArray("[[1,1]]")
	target = 2
	assertEq(false, searchMatrix(matrix, target))
}

func assertEq[T comparable](a, b T) {
	if a != b {
		fmt.Printf("Failed: %v != %v\n", a, b)
	} else {
		fmt.Printf("Ok: %v\n", a)
	}
}

func searchMatrix(matrix [][]int, target int) bool {
	ROW := len(matrix)
	COL := len(matrix[0])
	left := 0
	right := ROW*COL - 1

	for left <= right {
		mid := left + (right-left)/2
		midRow := mid / COL
		midCol := mid % COL
		if matrix[midRow][midCol] == target {
			return true
		} else if matrix[midRow][midCol] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return false
}

func parse2DArray(s string) [][]int {
	s = s[1 : len(s)-1]
	parts := strings.Split(s, "],[")
	var res [][]int
	for _, p := range parts {
		p = strings.Trim(p, "[]")
		ints := strings.Split(p, ",")
		var cur []int
		for _, i := range ints {
			n, err := strconv.ParseInt(i, 10, 0)
			if err != nil {
				panic(err)
			}
			cur = append(cur, int(n))
		}
		res = append(res, cur)
	}
	return res
}
