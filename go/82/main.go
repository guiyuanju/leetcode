package main

import "fmt"

func main() {
	printList(deleteDuplicates(makeList([]int{1, 2, 3, 3, 4, 4, 5})))
	printList(deleteDuplicates(makeList([]int{1, 1, 1, 2, 3})))
}

func deleteDuplicates(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != nil && cur.Next != nil {
		if cur.Val != cur.Next.Val {
			prev = cur
			cur = cur.Next
			continue
		}

		for cur.Next != nil && cur.Val == cur.Next.Val {
			cur = cur.Next
		}
		cur = cur.Next
		if prev == nil {
			head = cur
		} else {
			prev.Next = cur
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
