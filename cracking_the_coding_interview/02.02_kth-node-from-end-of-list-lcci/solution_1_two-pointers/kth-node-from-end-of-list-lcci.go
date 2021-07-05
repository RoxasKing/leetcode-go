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
func kthToLast(head *ListNode, k int) int {
	last := head
	for ; k > 1 && last != nil; k-- {
		last = last.Next
	}
	if last == nil {
		return -1
	}
	out := head
	for last.Next != nil {
		out = out.Next
		last = last.Next
	}
	return out.Val
}

type ListNode struct {
	Val  int
	Next *ListNode
}
