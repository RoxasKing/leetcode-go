package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
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
