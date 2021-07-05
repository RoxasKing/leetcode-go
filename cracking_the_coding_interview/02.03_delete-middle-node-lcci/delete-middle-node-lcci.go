package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(node *ListNode) {
	for {
		node.Val = node.Next.Val
		if node.Next.Next == nil {
			break
		}
		node = node.Next
	}
	node.Next = nil
}

type ListNode struct {
	Val  int
	Next *ListNode
}
