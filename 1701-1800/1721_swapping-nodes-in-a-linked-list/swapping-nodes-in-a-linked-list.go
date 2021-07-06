package main

// Tags:
// Two Pointers

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapNodes(head *ListNode, k int) *ListNode {
	headPre := &ListNode{Next: head}
	ptr1 := headPre
	last := headPre
	for i := 0; i < k; i++ {
		ptr1 = ptr1.Next
		last = last.Next
	}
	ptr2 := headPre
	for last != nil {
		ptr2 = ptr2.Next
		last = last.Next
	}
	ptr1.Val, ptr2.Val = ptr2.Val, ptr1.Val
	return headPre.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
