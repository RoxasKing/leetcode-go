package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	headPre := &ListNode{Next: head}
	ptr := headPre
	for ptr.Next != nil {
		if ptr.Next.Val == val {
			ptr.Next = ptr.Next.Next
		} else {
			ptr = ptr.Next
		}
	}
	return headPre.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
