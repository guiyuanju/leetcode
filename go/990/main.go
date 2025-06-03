package main

import "fmt"

func main() {
	equations := []string{"a==b", "b!=a"}
	assertEq(false, equationsPossible(equations))

	equations = []string{"b==a", "a==b"}
	assertEq(true, equationsPossible(equations))

	equations = []string{"a==b", "b!=c", "c==a"}
	assertEq(false, equationsPossible(equations))

	equations = []string{"a==b", "b==c", "d==e", "e==f", "d==a", "f!=a"}
	assertEq(false, equationsPossible(equations))
}

func assertEq[T comparable](a, b T) {
	if a != b {
		fmt.Printf("Failed: %v != %v\n", a, b)
	} else {
		fmt.Printf("Ok: %v\n", a)
	}
}

func equationsPossible(equations []string) bool {
	uf := make([]int, 26)
	for i := range uf {
		uf[i] = i
	}
	for _, e := range equations {
		l := int(e[0] - 'a')
		r := int(e[3] - 'a')
		isEqual := e[1] != '!'
		if !isEqual {
			continue
		}
		lGroupID := uf[l]
		for i := range uf {
			if uf[i] == lGroupID {
				uf[i] = uf[r]
			}
		}
	}

	for _, e := range equations {
		l := int(e[0] - 'a')
		r := int(e[3] - 'a')
		isEqual := e[1] != '!'
		if isEqual {
			continue
		}
		if uf[l] == uf[r] {
			return false
		}
	}

	return true
}
