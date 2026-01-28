package main

import (
	"fmt"
	"reflect"
)

func main() {
	assertEq("Radiant", predictPartyVictory("RD"))
	assertEq("Dire", predictPartyVictory("RDD"))
}

func assertEq(a, b any) {
	if reflect.DeepEqual(a, b) {
		fmt.Printf("Ok: %v\n", a)
	} else {
		fmt.Printf("Failed: %v != %v\n", a, b)
	}
}

func predictPartyVictory(senate string) string {
	var r, d []int
	for i, v := range []byte(senate) {
		if v == 'R' {
			r = append(r, i)
		} else {
			d = append(d, i)
		}
	}

	next := len(senate)
	for len(r) > 0 && len(d) > 0 {
		if r[0] < d[0] {
			r = append(r, r[0]+next)
		} else {
			d = append(d, d[0]+next)
		}
		d = d[1:]
		r = r[1:]
		next++
	}

	if len(r) == 0 {
		return "Dire"
	}
	return "Radiant"
}
