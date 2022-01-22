package main

// Difficulty:
// Medium

// Tags:
// Two Pointers

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		if slow, fast = slow.Next, fast.Next.Next; slow == fast {
			break
		}
	}
	if fast == nil || fast.Next == nil {
		return nil
	}
	for slow = head; slow != fast; {
		slow, fast = slow.Next, fast.Next
	}
	return slow
}
