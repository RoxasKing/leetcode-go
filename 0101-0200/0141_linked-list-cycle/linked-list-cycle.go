package main

// Difficulty:
// Easy

// Tags:
// Two Pointers

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	for slow, fast := head, head.Next; fast != nil && fast.Next != nil; {
		if slow, fast = slow.Next, fast.Next.Next; slow == fast {
			return true
		}
	}
	return false
}
