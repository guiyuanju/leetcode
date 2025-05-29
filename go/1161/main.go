package main

import "fmt"
import "math"

func main() {
	root := makeBinaryTree([]any{1, 7, 0, 7, -8, nil, nil})
	assertEq(2, maxLevelSum(root))

	root = makeBinaryTree([]any{989, nil, 10250, 98693, -89388, nil, nil, nil, -32127})
	assertEq(2, maxLevelSum(root))

	root = makeBinaryTree([]any{-100, -200, -300, -20, -5, -10, nil})
	assertEq(3, maxLevelSum(root))
}

func maxLevelSum(root *TreeNode) int {
	var level, curLevel int
	sum := math.MinInt
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		length := len(queue)
		curSum := 0
		curLevel++
		for range length {
			cur := queue[0]
			queue = queue[1:]
			curSum += cur.Val
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		if curSum > sum {
			sum = curSum
			level = curLevel
		}
	}
	return level
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func makeBinaryTree(nodes []any) *TreeNode {
	var root *TreeNode
	queue := []*TreeNode{}
	for i, n := range nodes {
		var cur *TreeNode
		if n != nil {
			cur = &TreeNode{
				Val:   n.(int),
				Left:  nil,
				Right: nil,
			}
		}
		queue = append(queue, cur)

		for queue[0] == nil {
			queue = queue[1:]
		}

		if i == 0 {
			root = cur
		} else if i%2 != 0 {
			queue[0].Left = cur
		} else {
			queue[0].Right = cur
			queue = queue[1:]
		}
	}
	return root
}

func printBinaryTree(root *TreeNode) {
	flat := []*TreeNode{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		flat = append(flat, cur)
		if cur != nil {
			queue = append(queue, cur.Left)
			queue = append(queue, cur.Right)
		}
	}
	for i := len(flat) - 1; i >= 0 && flat[i] == nil; i-- {
		flat = flat[:len(flat)-1]
	}
	fmt.Print("[")
	for i, v := range flat {
		if v != nil {
			fmt.Print(v.Val)
		} else {
			fmt.Print("null")
		}
		if i < len(flat)-1 {
			fmt.Print(", ")
		}
	}
	fmt.Println("]")
}

func assertEq[T comparable](a, b T) {
	if a != b {
		fmt.Printf("Failed: %v != %v\n", a, b)
	} else {
		fmt.Printf("Ok: %v\n", a)
	}
}
