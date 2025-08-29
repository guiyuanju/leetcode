package main

import "fmt"

func main() {
	fmt.Println(predictPartyVictory("RD"))
	fmt.Println(predictPartyVictory("RDD"))
}

const R byte = 'R'
const D byte = 'D'

func predictPartyVictory(senate string) string {
	rs := []int{}
	ds := []int{}
	for i, c := range senate {
		if byte(c) == R {
			rs = append(rs, i)
		} else {
			ds = append(ds, i)
		}
	}

	idx := len(senate)
	for len(rs) > 0 && len(ds) > 0 {
		if rs[0] < ds[0] {
			rs = append(rs, idx)
		} else {
			ds = append(ds, idx)
		}
		rs = rs[1:]
		ds = ds[1:]
		idx++
	}

	if len(rs) == 0 {
		return "Dire"
	}
	return "Radiant"
}
