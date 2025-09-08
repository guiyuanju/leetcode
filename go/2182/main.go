package main

import "fmt"

func main() {
	s := "cczazcc"
	repeatLimit := 3
	assertEq("zzcccac", repeatLimitedString(s, repeatLimit))

	s = "aababab"
	repeatLimit = 2
	assertEq("bbabaa", repeatLimitedString(s, repeatLimit))

	s = "bplpcfifosybmjxphbxdltxtfrjspgixoxzbpwrtkopepjxfooazjyosengdlvyfchqhqxznnhuuxhtbrojyhxwlsrklsryvmufoibgfyxgjw"
	repeatLimit = 1
	assertEq("zyzyzyxyxyxyxwxwxwxvxvxuxututststsrsrsrqrqrpopopopopopopononmnmlklkljljljijijijhghghghghfhfefefdfdfcfcbab", repeatLimitedString(s, repeatLimit))
}

func assertEq[T comparable](a, b T) {
	if a != b {
		fmt.Printf("Failed: %v != %v\n", a, b)
	} else {
		fmt.Printf("Ok: %v\n", a)
	}
}

func repeatLimitedString(s string, repeatLimit int) string {
	count := [26]int{}
	for _, c := range []byte(s) {
		count[c-'a']++
	}

	var res []byte
	for j := 25; j >= 0; {
		if count[j] == 0 {
			j--
			continue
		}
		for range min(count[j], repeatLimit) {
			res = append(res, 'a'+byte(j))
			count[j]--
		}
		if count[j] > 0 {
			i := j - 1
			for ; i >= 0; i-- {
				if count[i] != 0 {
					res = append(res, 'a'+byte(i))
					count[i]--
					break
				}
			}
			if i < 0 {
				count[j] = 0
			}
		} else {
			j--
		}
	}
	return string(res)
}
