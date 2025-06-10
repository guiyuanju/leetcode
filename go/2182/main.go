package main

import "fmt"
import "container/heap"

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
	count := map[byte]int{}
	for _, r := range s {
		c := byte(r)
		count[c]++
	}

	h := new(Heap)
	for k, v := range count {
		heap.Push(h, &Item{k, v})
	}

	res := make([]byte, 0, len(s))
	for h.Len() > 0 {
		cur := heap.Pop(h).(*Item)
		repeat := min(cur.count, repeatLimit)
		for range repeat {
			res = append(res, cur.value)
		}

		if cur.count > repeat && h.Len() > 0 {
			next := heap.Pop(h).(*Item)
			res = append(res, next.value)
			next.count--
			if next.count > 0 {
				heap.Push(h, next)
			}
			cur.count -= repeat
			heap.Push(h, cur)
		}
	}

	return string(res)
}

type Heap []*Item
type Item struct {
	value byte
	count int
}

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i].value > h[j].value }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x any)        { *h = append(*h, x.(*Item)) }
func (h *Heap) Pop() any {
	old := *h
	n := len(old)
	res := old[n-1]
	*h = old[:n-1]
	return res
}
