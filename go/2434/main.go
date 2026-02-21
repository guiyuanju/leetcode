package main

import (
	"fmt"
	"reflect"
)

func main() {
	assertEq("azz", robotWithString("zza"))
	assertEq("abc", robotWithString("bac"))
	assertEq("addb", robotWithString("bdda"))
	assertEq("bdevfziy", robotWithString("bydizfve"))
}

func robotWithString(s string) string {
	t := []byte{}
	res := []byte{}

	mins := make([]byte, len(s))
	mins[len(s)-1] = s[len(s)-1]
	for i := len(s) - 2; i >= 0; i-- {
		mins[i] = min(s[i], mins[i+1])
	}

	for i := 0; i < len(s); {
		if len(t) == 0 || mins[i] < t[len(t)-1] {
			t = append(t, s[i])
			i++
		} else {
			res = append(res, t[len(t)-1])
			t = t[:len(t)-1]
		}
	}

	for i := len(t) - 1; i >= 0; i-- {
		res = append(res, t[i])
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
