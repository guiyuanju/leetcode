package main

import (
	"fmt"
	"reflect"
	"slices"
	"time"
)

func main() {
	assertEq(true, canPartitionKSubsets([]int{4, 3, 2, 3, 5, 2, 1}, 4))
	assertEq(false, canPartitionKSubsets([]int{1, 2, 3, 4}, 3))
	assertEq(false, canPartitionKSubsets([]int{2, 2, 2, 2, 3, 4, 5}, 4))
	assertEq(true, canPartitionKSubsets([]int{18, 20, 39, 73, 96, 99, 101, 111, 114, 190, 207, 295, 471, 649, 700, 1037}, 4))
}

func assertEq(a, b any) {
	if reflect.DeepEqual(a, b) {
		fmt.Printf("Ok: %v\n", a)
	} else {
		fmt.Printf("Failed: %v != %v\n", a, b)
	}
}

func canPartitionKSubsets(nums []int, k int) bool {
	now := time.Now()
	res := canPartitionKSubsets_dp(nums, k)
	fmt.Println("dp:", time.Since(now))

	now = time.Now()
	res = canPartitionKSubsets_bt(nums, k)
	fmt.Println("bt:", time.Since(now))

	return res
}

func canPartitionKSubsets_dp(nums []int, k int) bool {
	var sum int
	for _, n := range nums {
		sum += n
	}
	if sum%k > 0 {
		return false
	}
	length := sum / k

	type state struct {
		i    int
		conf [16]int
	}

	memo := map[state]bool{}

	slices.SortFunc(nums, func(a, b int) int { return b - a })

	var dp func(i int, conf [16]int) bool
	dp = func(i int, conf [16]int) bool {
		if i >= len(nums) {
			return true
		}

		slices.Sort(conf[:k])

		if v, ok := memo[state{i, conf}]; ok {
			return v
		}

		for j := range k {
			if conf[j]+nums[i] <= length {
				conf[j] += nums[i]
				res := dp(i+1, conf)
				conf[j] -= nums[i]
				if res {
					memo[state{i, conf}] = true
					return true
				}
			}
		}

		memo[state{i, conf}] = false
		return false
	}

	return dp(0, [16]int{})
}

func canPartitionKSubsets_bt(nums []int, k int) bool {
	var sum int
	for _, n := range nums {
		sum += n
	}
	if sum%k > 0 {
		return false
	}
	length := sum / k

	ks := make([]int, k)

	var bt func(i int) bool
	bt = func(i int) bool {
		if i >= len(nums) {
			return true
		}

		for j := range k {
			if ks[j]+nums[i] <= length {
				ks[j] += nums[i]
				if bt(i + 1) {
					return true
				}
				ks[j] -= nums[i]
			}
		}

		return false
	}

	return bt(0)
}
