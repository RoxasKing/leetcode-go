package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	headPre := &ListNode{Next: head}
	for ptr := headPre; ptr.Next != nil; {
		if ptr.Next.Val < x && ptr != headPre {
			node := ptr.Next
			ptr.Next = ptr.Next.Next
			node.Next = headPre.Next
			headPre.Next = node
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
