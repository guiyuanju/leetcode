package main

import "fmt"

func main() {
	printList(deleteDuplicates(makeList([]int{1, 1, 2})))
	printList(deleteDuplicates(makeList([]int{1, 1, 2, 3, 3})))
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	cur := head
	next := head.Next
	for next != nil {
		if cur.Val == next.Val {
			cur.Next = next.Next
			next = next.Next
		} else {
			next = next.Next
			cur = cur.Next
		}
	}

	return head
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
