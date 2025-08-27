package main

import "fmt"

func main() {
	fmt.Println(maximumRobots([]int{3, 6, 1, 3, 4}, []int{2, 1, 3, 4, 5}, 25))
	fmt.Println(maximumRobots([]int{11, 12, 19}, []int{10, 8, 7}, 19))
}

func maximumRobots(chargeTimes []int, runningCosts []int, budget int64) int {
	mono := []int{}
	var i, sum, res int
	for j := range chargeTimes {
		sum += runningCosts[j]

		for len(mono) > 0 && chargeTimes[j] > mono[len(mono)-1] {
			mono = mono[:len(mono)-1]
		}
		mono = append(mono, chargeTimes[j])

		for len(mono) > 0 && int64(mono[0]+(j-i+1)*sum) > budget {
			sum -= runningCosts[i]
			if chargeTimes[i] == mono[0] {
				mono = mono[1:]
			}
			i++
		}

		res = max(res, j-i+1)
	}

	return res
}
