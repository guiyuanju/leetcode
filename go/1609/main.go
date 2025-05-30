package main

import "fmt"

func main() {
	root := makeBinaryTree([]any{1, 10, 4, 3, nil, 7, 9, 12, 8, 6, nil, nil, 2})
	assertEq(true, isEvenOddTree(root))

	root = makeBinaryTree([]any{5, 4, 2, 3, 3, 7})
	assertEq(false, isEvenOddTree(root))

	root = makeBinaryTree([]any{5, 9, 1, 3, 5, 7})
	assertEq(false, isEvenOddTree(root))
}

func assertEq[T comparable](a, b T) {
	if a != b {
		fmt.Printf("Failed: %v != %v\n", a, b)
	} else {
		fmt.Printf("Ok: %v\n", a)
	}
}

func isEvenOddTree(root *TreeNode) bool {
	isEven := true
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		length := len(queue)
		if (isEven && increasing(queue[:length])) || (!isEven && decreasing(queue[:length])) {
			for range length {
				cur := queue[0]
				queue = queue[1:]
				if isEven && cur.Val%2 == 0 {
					return false
				}
				if !isEven && cur.Val%2 != 0 {
					return false
				}
				if cur.Left != nil {
					queue = append(queue, cur.Left)
				}
				if cur.Right != nil {
					queue = append(queue, cur.Right)
				}
			}
		} else {
			return false
		}

		isEven = !isEven
	}

	return true
}

func increasing(queue []*TreeNode) bool {
	if len(queue) <= 1 {
		return true
	}
	prev := queue[0]
	for i := 1; i < len(queue); i++ {
		if queue[i].Val <= prev.Val {
			return false
		}
		prev = queue[i]
	}
	return true
}

func decreasing(queue []*TreeNode) bool {
	if len(queue) <= 1 {
		return true
	}
	prev := queue[0]
	for i := 1; i < len(queue); i++ {
		if queue[i].Val >= prev.Val {
			return false
		}
		prev = queue[i]
	}
	return true
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
