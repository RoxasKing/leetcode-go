package main

// Difficulty:
// Medium

// Tags:
// Two Pointers

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	n := 1
	tail := head
	for ; tail.Next != nil; tail = tail.Next {
		n++
	}
	if k = n - k%n; k == 0 || k == n {
		return head
	}
	for ; k > 0; k-- {
		next := head.Next
		tail.Next = head
		tail = head
		head = next
	}
	tail.Next = nil
	return head
}
