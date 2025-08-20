package main

import "fmt"

func main() {
	printList(reverseBetween(makeList([]int{1, 2, 3, 4, 5}), 2, 4))
	printList(reverseBetween(makeList([]int{5}), 1, 1))
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	var start, end, first, last *ListNode

	// find start
	i := 1
	start = head
	for i < left {
		first = start
		start = start.Next
		i++
	}

	// find last
	end = start
	for ; i < right; i++ {
		end = end.Next
	}
	last = end.Next

	// reverse
	var prev *ListNode
	cur := start
	for prev != end {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	// deal with two ends
	if first != nil {
		first.Next = end
		start.Next = last
		return head
	}
	start.Next = last
	return end
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func makeList(vals []int) *ListNode {
	head := &ListNode{0, nil}
	prev := head
	for _, v := range vals {
		prev.Next = &ListNode{v, nil}
		prev = prev.Next
	}
	return head.Next
}

func printList(head *ListNode) {
	fmt.Print("[")
	for head != nil {
		fmt.Printf("%d, ", head.Val)
		head = head.Next
	}
	fmt.Println("]")
}
