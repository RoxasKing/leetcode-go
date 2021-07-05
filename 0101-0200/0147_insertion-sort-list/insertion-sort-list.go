package main

// Tags:
// Insertion Sort + Linked List
func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	headPre := &ListNode{Next: head}
	last := head
	ptr := head.Next
	head.Next = nil
	for ptr != nil {
		ptrNext := ptr.Next
		if ptr.Val > last.Val {
			last.Next = ptr
			last = last.Next
			last.Next = nil
		} else {
			n := headPre
			for n.Next != nil && n.Next.Val <= ptr.Val {
				n = n.Next
			}
			nNext := n.Next
			n.Next = ptr
			ptr.Next = nNext
		}
		ptr = ptrNext
	}
	return headPre.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
