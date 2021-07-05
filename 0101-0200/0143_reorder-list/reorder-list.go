package main

// Tags:
// Two Pointers

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	l1, l2 := head, slow.Next
	slow.Next = nil
	l2 = reverse(l2)
	for l1 != nil && l2 != nil {
		next1 := l1.Next
		next2 := l2.Next
		l1.Next = l2
		l2.Next = next1
		l1, l2 = next1, next2
	}
}

func reverse(head *ListNode) *ListNode {
	var out *ListNode
	for head != nil {
		next := head.Next
		head.Next = out
		out = head
		head = next
	}
	return out
}

type ListNode struct {
	Val  int
	Next *ListNode
}
