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
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}

	cnt := 0
	for n := head; n != nil; n = n.Next {
		cnt++
	}

	k %= cnt
	if k == 0 {
		return head
	}

	last := head
	for k > 0 {
		last = last.Next
		k--
	}

	ptr := head
	for last.Next != nil {
		ptr = ptr.Next
		last = last.Next
	}

	newHead := ptr.Next
	ptr.Next = nil
	last.Next = head

	return newHead
}

type ListNode struct {
	Val  int
	Next *ListNode
}
