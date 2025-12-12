package main

import (
	"fmt"
	"reflect"
	"slices"
)

func main() {
	assertEq("azz", robotWithString("zza"))
	assertEq("abc", robotWithString("bac"))
	assertEq("addb", robotWithString("bdda"))
	assertEq("bdevfziy", robotWithString("bydizfve"))
}

func robotWithString(s string) string {
	mins := make([]byte, len(s))
	mins[len(mins)-1] = s[len(s)-1]
	for i := len(s) - 2; i >= 0; i-- {
		mins[i] = min(mins[i+1], s[i])
	}

	var t, res []byte
	var i int
	for i < len(s) {
		if len(t) > 0 && t[len(t)-1] <= mins[i] {
			res = append(res, t[len(t)-1])
			t = t[:len(t)-1]
		} else {
			t = append(t, s[i])
			i++
		}
	}

	if len(t) > 0 {
		slices.Reverse(t)
		res = append(res, t...)
	}

	return string(res)
}

func assertEq(a, b any) {
	if reflect.DeepEqual(a, b) {
		fmt.Printf("Ok: %v\n", a)
	} else {
		fmt.Printf("Failed: %v != %v\n", a, b)
	}
}
