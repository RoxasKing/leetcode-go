package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	ptr := head
	for node := ptr.Next; node != nil; node = node.Next {
		if node.Val != ptr.Val {
			ptr.Next = node
			ptr = ptr.Next
		}
	}
	ptr.Next = nil
	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}
