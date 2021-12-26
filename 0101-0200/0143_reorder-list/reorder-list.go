package main

// Difficulty:
// Medium

// Tags:
// Two Pointers

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	l1, l2 := head, reverse(slow.Next)
	slow.Next = nil
	for l2 != nil {
		t1, t2 := l1.Next, l2.Next
		l1.Next = l2
		l2.Next = t1
		l1, l2 = t1, t2
	}
}

func reverse(head *ListNode) *ListNode {
	var out *ListNode
	for head != nil {
		next := head.Next
		head.Next = out
		out, head = head, next
	}
	return out
}
