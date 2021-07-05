package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getKthFromEnd(head *ListNode, k int) *ListNode {
	last := head
	for ; k > 0; k-- {
		last = last.Next
	}
	for last != nil {
		head = head.Next
		last = last.Next
	}
	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}
