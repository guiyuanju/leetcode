package main

import "fmt"

func main() {
	printList(swapPairs(makeList([]int{1, 2, 3, 4})))
	printList(swapPairs(makeList([]int{})))
	printList(swapPairs(makeList([]int{1})))
	printList(swapPairs(makeList([]int{1, 2, 3})))
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var prev *ListNode
	cur := head
	next := head.Next
	dump := next
	for cur != nil && next != nil {
		cur.Next = next.Next
		next.Next = cur
		if prev != nil {
			prev.Next = next
		}
		prev = cur
		cur = cur.Next
		if cur != nil {
			next = cur.Next
		}
	}

	return dump
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
