package main

// Difficulty:
// Medium

// Tags:
// Linked List

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteMiddle(head *ListNode) *ListNode {
	n := 0
	for p := head; p != nil; p = p.Next {
		n++
	}
	if n == 1 {
		return nil
	}
	p := head
	for i := 0; i < n/2-1; i++ {
		p = p.Next
	}
	p.Next = p.Next.Next
	return head
}
