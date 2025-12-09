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
}

// Can't directly use Dijkstra algo there,
// becuase only the lowest price path is computed, the steps is neglected,
// the path may end up with more than k+1 steps,
// we can use BFS which proceed one level a time,
// and recompute the same node if it has been visited before,
// in this way, we can iterate all paths to the dest node,
// by limiting the steps to k+1, we not only compute k stops merely,
// but also prevent infinite loop in the case of a loop in graph.
//
// There exists other solutions, such as Dijkstra algo:
// normally, Dijkstra algo tracks the price to d node, but we track stop
// instead, only traverse an edgo to a node if it has not alreadly been visited
// with fewer stops, which enable it visiting a node with higher price
// but have fewer stops (more stops is impossible to be the answer),
//
// Bellman Ford algotithm is also possible.
func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	graph := make(map[int][][2]int, n)
	for _, f := range flights {
		graph[f[0]] = append(graph[f[0]], [2]int{f[1], f[2]})
	}

	type step struct {
		node, price int
	}

	prices := make([]int, n)
	for i := range prices {
		prices[i] = math.MaxInt
	}

	// We use BFS approach, with a optimization: prices,
	// which prevent revisiting a node if it already visited with a lower prices,
	// which is optional. But required for leetcode LTE.
	var stop int
	queue := []step{{src, 0}}
	for stop <= k && len(queue) > 0 {
		for range len(queue) {
			cur := queue[0]
			queue = queue[1:]
			for _, nei := range graph[cur.node] {
				newPrice := cur.price + nei[1]
				if newPrice < prices[nei[0]] {
					prices[nei[0]] = newPrice
					queue = append(queue, step{nei[0], newPrice})
				}
			}
		}
		stop++
	}

	if prices[dst] == math.MaxInt {
		return -1
	}
	return prices[dst]
}
