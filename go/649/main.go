package main

import "fmt"

func main() {
	fmt.Println(predictPartyVictory("RD"))
	fmt.Println(predictPartyVictory("RDD"))
}

const (
	R byte = 'R'
	D byte = 'D'
)

func predictPartyVictory(senate string) string {
	var rq, dq []int
	for i, s := range []byte(senate) {
		if s == R {
			rq = append(rq, i)
		} else {
			dq = append(dq, i)
		}
	}

	for nextIdx := len(senate); len(rq) > 0 && len(dq) > 0; nextIdx++ {
		if rq[0] < dq[0] {
			rq = append(rq, nextIdx)
		} else {
			dq = append(dq, nextIdx)
		}
		rq = rq[1:]
		dq = dq[1:]
	}

	if len(rq) > 0 {
		return "Radiant"
	}
	return "Dire"
}
