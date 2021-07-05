package main

// Tags:
// Linked List
func partition(head *ListNode, x int) *ListNode {
	leftHead := &ListNode{Val: -1}
	leftTail := leftHead
	rightHead := &ListNode{Val: -1}
	rightTail := rightHead
	for n := head; n != nil; n = n.Next {
		if n.Val < x {
			leftTail.Next = n
			leftTail = leftTail.Next
		} else {
			rightTail.Next = n
			rightTail = rightTail.Next
		}
	}
	leftTail.Next = rightHead.Next
	rightTail.Next = nil
	return leftHead.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
