package main

import (
	"fmt"
	"slices"
)

func main() {
	products := []string{"mobile", "mouse", "moneypot", "monitor", "mousepad"}
	searchWrod := "mouse"
	fmt.Println(suggestedProducts(products, searchWrod))
}

func suggestedProducts(products []string, searchWord string) [][]string {
	type Node struct {
		val      byte
		children map[byte]*Node
		words    []string
	}
	newNode := func(val byte) *Node {
		return &Node{val, map[byte]*Node{}, nil}
	}

	root := newNode(0)

	for _, p := range products {
		cur := root
		for _, c := range p {
			if _, ok := cur.children[byte(c)]; !ok {
				cur.children[byte(c)] = newNode(byte(c))
			}
			cur = cur.children[byte(c)]
			cur.words = append(cur.words, p)
			slices.Sort(cur.words)
			if len(cur.words) > 3 {
				cur.words = cur.words[:3]
			}
		}
	}

	var res [][]string
	cur := root
	for _, c := range []byte(searchWord) {
		if v, ok := cur.children[c]; ok {
			cur = v
			res = append(res, cur.words)
		} else {
			// make all following loop add nil
			cur.children = map[byte]*Node{}
			res = append(res, nil)
		}
	}

	return res
}
