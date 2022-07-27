package main

// Difficulty:
// Medium

// Tags:
// Linked List

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}
	var a, b, atail, btail *ListNode
	for p := head; p != nil; p = p.Next {
		if p.Val < x {
			if a == nil {
				a, atail = p, p
			} else {
				atail.Next = p
				atail = p
			}
		} else {
			if b == nil {
				b, btail = p, p
			} else {
				btail.Next = p
				btail = p
			}
		}
	}
	if a == nil || b == nil {
		return head
	}
	atail.Next = b
	btail.Next = nil
	return a
}
