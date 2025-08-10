package main

import (
	"math"

	"github.com/guiyuanju/lcutils/assert"
)

func main() {
	assert.Eq(700, findCheapestPrice(4, [][]int{{0, 1, 100}, {1, 2, 100}, {2, 0, 100}, {1, 3, 600}, {2, 3, 200}}, 0, 3, 1))
	assert.Eq(200, findCheapestPrice(3, [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}, 0, 2, 1))
	assert.Eq(500, findCheapestPrice(3, [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}, 0, 2, 0))
	assert.Eq(7, findCheapestPrice(5, [][]int{{0, 1, 5}, {1, 2, 5}, {0, 3, 2}, {3, 1, 2}, {1, 4, 1}, {4, 2, 1}}, 0, 2, 2))
}

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	g := map[int][][]int{}
	for _, f := range flights {
		g[f[0]] = append(g[f[0]], []int{f[1], f[2]})
	}

	prices := make([]int, n)
	for i := range prices {
		prices[i] = math.MaxInt
	}

	queue := []Node{{src, 0, 0}}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if cur.price >= prices[cur.city] {
			continue
		}
		prices[cur.city] = cur.price
		stop := cur.stop + 1
		if cur.city == src {
			stop = 0
		}
		if stop > k {
			continue
		}
		for _, nei := range g[cur.city] {
			price := cur.price + nei[1]
			if price < prices[nei[0]] {
				queue = append(queue, Node{nei[0], price, stop})
			}
		}
	}

	if prices[dst] == math.MaxInt {
		return -1
	}
	return prices[dst]
}

type Node struct {
	city, price, stop int
}
