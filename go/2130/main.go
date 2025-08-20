package main

import "fmt"

func main() {
	fmt.Println(pairSum(makeList([]int{5, 4, 2, 1})))
	fmt.Println(pairSum(makeList([]int{4, 2, 2, 3})))
	fmt.Println(pairSum(makeList([]int{1, 100000})))
}

func pairSum(head *ListNode) int {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	var prev *ListNode
	cur := slow
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	head1 := head
	head2 := prev
	var res int
	for head2 != nil {
		res = max(res, head1.Val+head2.Val)
		head1 = head1.Next
		head2 = head2.Next
	}

	return res
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
