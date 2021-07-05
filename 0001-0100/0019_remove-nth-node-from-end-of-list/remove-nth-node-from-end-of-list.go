package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	headPre := &ListNode{Next: head}
	last := headPre
	for n > 0 {
		last = last.Next
		n--
	}
	ptr := headPre
	for last.Next != nil {
		last = last.Next
		ptr = ptr.Next
	}
	ptr.Next = ptr.Next.Next
	return headPre.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
