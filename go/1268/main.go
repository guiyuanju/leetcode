package main

import (
	"fmt"
	"slices"
)

func main() {
	products := []string{"mobile", "mouse", "moneypot", "monitor", "mousepad"}
	searchWrod := "mouse"
	fmt.Println(suggestedProducts(products, searchWrod))

	fmt.Println(suggestedProducts([]string{"havana"}, "havana"))
}

func suggestedProducts(products []string, searchWord string) [][]string {
	type Node struct {
		children map[byte]*Node
		strings  []string
	}

	slices.Sort(products)

	root := Node{map[byte]*Node{}, []string{}}
	for _, p := range products {
		cur := &root
		for _, c := range []byte(p) {
			if _, ok := cur.children[c]; !ok {
				cur.children[c] = &Node{map[byte]*Node{}, []string{}}
			}
			cur = cur.children[c]
			if len(cur.strings) < 3 {
				cur.strings = append(cur.strings, p)
			}
		}
	}

	var res [][]string
	cur := &root
	for _, c := range []byte(searchWord) {
		if _, ok := cur.children[c]; ok {
			cur = cur.children[c]
			res = append(res, cur.strings)
		} else {
			cur.children = map[byte]*Node{}
			res = append(res, []string{})
		}
	}
	return res
}
