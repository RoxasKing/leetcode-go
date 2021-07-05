package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	headPre := &ListNode{Next: head}
	ptr := headPre
	for ptr.Next != nil && ptr.Next.Next != nil {
		next := ptr.Next.Next.Next
		a, b := ptr.Next, ptr.Next.Next
		ptr.Next = b
		b.Next = a
		a.Next = next
		ptr = a
	}
	return headPre.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
