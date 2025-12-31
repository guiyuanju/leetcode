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
	assert.Eq(-1, findCheapestPrice(3, [][]int{{4, 1, 1}, {1, 2, 3}, {0, 3, 2}, {0, 4, 10}, {3, 1, 1}, {1, 4, 3}}, 2, 1, 1))
	assert.Eq(6, findCheapestPrice(4, [][]int{{0, 1, 1}, {0, 2, 5}, {1, 2, 1}, {2, 3, 1}}, 0, 3, 1))
	assert.Eq(7, findCheapestPrice(5, [][]int{{0, 1, 1}, {0, 2, 5}, {1, 2, 1}, {2, 3, 1}, {3, 4, 1}}, 0, 4, 2))
}

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	g := make(map[int][][2]int, n)
	for _, f := range flights {
		g[f[0]] = append(g[f[0]], [2]int{f[1], f[2]})
	}

	// 用于防止不必要的访问
	// 1. 更大的k, 但是更大的price
	// 2. 情况一同时包含环路的防止
	prices := make([]int, n)
	for i := range prices {
		prices[i] = math.MaxInt
	}
	prices[src] = 0

	// 每一步都需要单独的price
	// 同一个节点可能有小k大price和大k小price两种
	// 不能相互覆盖
	type ele struct {
		node, price int
	}

	q := []ele{{src, 0}}
	var step int
	for ; len(q) > 0 && step <= k; step++ {
		for range len(q) {
			cur := q[0]
			q = q[1:]
			for _, nei := range g[cur.node] {
				if cur.price+nei[1] < prices[nei[0]] {
					prices[nei[0]] = cur.price + nei[1]
					q = append(q, ele{nei[0], prices[nei[0]]})
				}
			}
		}
	}

	if prices[dst] == math.MaxInt {
		return -1
	}
	return prices[dst]
}
