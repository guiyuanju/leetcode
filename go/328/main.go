package main

import "fmt"

func main() {
	printList(oddEvenList(makeList([]int{1, 2, 3, 4, 5})))
	printList(oddEvenList(makeList([]int{2, 1, 3, 5, 6, 4, 7})))
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	odd := head
	even := head.Next
	dump := even

	for odd.Next != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}

	if even != nil {
		even.Next = nil
	}
	odd.Next = dump

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
