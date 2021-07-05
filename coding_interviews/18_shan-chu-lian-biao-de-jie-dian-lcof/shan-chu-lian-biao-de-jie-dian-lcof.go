package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(head *ListNode, val int) *ListNode {
	pre := &ListNode{Next: head}
	for n := pre; n.Next != nil; n = n.Next {
		if n.Next.Val == val {
			n.Next = n.Next.Next
			break
		}
	}
	return pre.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
