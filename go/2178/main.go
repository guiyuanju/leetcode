package main

import "fmt"

func main() {
	fmt.Println(maximumEvenSplit(12))
	fmt.Println(maximumEvenSplit(7))
	fmt.Println(maximumEvenSplit(28))
}

func maximumEvenSplit(finalSum int64) []int64 {
	if finalSum%2 > 0 {
		return []int64{}
	}
	var res []int64
	var cur int64 = 2
	for finalSum >= 2*cur+2 {
		res = append(res, cur)
		finalSum -= cur
		cur += 2
	}
	if finalSum > 0 {
		res = append(res, finalSum)
	}
	return res
}
