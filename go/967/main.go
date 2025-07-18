package main

import "github.com/guiyuanju/lcutils/assert"

func main() {
	assert.Eq([]int{181, 292, 707, 818, 929}, numsSameConsecDiff(3, 7))
	assert.Eq([]int{10, 12, 21, 23, 32, 34, 43, 45, 54, 56, 65, 67, 76, 78, 87, 89, 98}, numsSameConsecDiff(2, 1))
	assert.Eq([]int{11, 22, 33, 44, 55, 66, 77, 88, 99}, numsSameConsecDiff(2, 0))
}

func numsSameConsecDiff(n int, k int) []int {
	var res []int

	var bt func(cur []int)
	bt = func(cur []int) {
		if len(cur) == n {
			var num int
			for _, v := range cur {
				num = num*10 + v
			}
			res = append(res, num)
			return
		}

		if len(cur) == 0 {
			for i := 1; i < 10; i++ {
				bt(append(cur, i))
			}
		} else {
			last := cur[len(cur)-1]
			nums := []int{last - k, last + k}
			if k == 0 {
				nums = []int{last}
			}
			for _, num := range nums {
				if 0 <= num && num <= 9 {
					bt(append(cur, num))
				}
			}
		}
	}

	bt(nil)

	return res
}
