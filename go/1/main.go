package main

import "sort"

func twoSum1(nums []int, target int) []int {
	exist := map[int]int{}
	for i, n := range nums {
		if j, ok := exist[target-n]; ok {
			return []int{i, j}
		}
		exist[n] = i
	}
	return nil
}
func twoSum2(nums []int, target int) []int {
	newNums := make([]item, len(nums))
	for i, n := range nums {
		newNums[i] = item{i, n}
	}
	sort.Slice(newNums, func(i, j int) bool {
		return newNums[i].val < newNums[j].val
	})
	for i, j := 0, len(newNums)-1; i < j; {
		t := newNums[i].val + newNums[j].val
		if t == target {
			return []int{newNums[i].idx, newNums[j].idx}
		} else if t > target {
			j--
		} else {
			i++
		}
	}
	return nil
}

type item struct {
	idx int
	val int
}
