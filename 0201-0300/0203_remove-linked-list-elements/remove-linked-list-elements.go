package main

// Difficulty:
// Easy

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	for ; head != nil && head.Val == val; head = head.Next {
	}
	if head == nil {
		return nil
	}
	for ptr := head; ptr.Next != nil; {
		if ptr.Next.Val == val {
			ptr.Next = ptr.Next.Next
		} else {
			ptr = ptr.Next
		}
	}
	return head
}
