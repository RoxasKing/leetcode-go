package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeDuplicateNodes(head *ListNode) *ListNode {
	headPre := &ListNode{Next: head}
	ptr := headPre
	mark := make([]bool, 20001)
	for node := head; node != nil; node = node.Next {
		if !mark[node.Val] {
			mark[node.Val] = true
			ptr.Next = node
			ptr = node
		}
	}
	ptr.Next = nil
	return headPre.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
