package main

import "fmt"

func main() {
	fmt.Println(findSmallestSetOfVertices(6, [][]int{{0, 1}, {0, 2}, {2, 5}, {3, 4}, {4, 2}}))
	fmt.Println(findSmallestSetOfVertices(5, [][]int{{0, 1}, {2, 1}, {3, 1}, {1, 4}, {2, 4}}))
}

func findSmallestSetOfVertices(n int, edges [][]int) []int {
	ind := make([]int, n)
	for _, e := range edges {
		ind[e[1]]++
	}

	var res []int
	for i, v := range ind {
		if v == 0 {
			res = append(res, i)
		}
	}

	return res
}
