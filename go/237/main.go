package main

import "fmt"

func main() {
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

func deleteNode(node *ListNode) {
	var prev *ListNode
	cur := node
	for cur.Next != nil {
		cur.Val = cur.Next.Val
		prev = cur
		cur = cur.Next
	}
	prev.Next = nil
}
