package main

// Difficulty:
// Medium

// Tags:
// Linked List

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	headPre := &ListNode{Next: head}
	pre := headPre
	for i := 1; i < left; i++ {
		pre = pre.Next
	}
	ptr := pre.Next
	for i := left + 1; i <= right; i++ {
		tmp := ptr.Next
		ptr.Next = tmp.Next
		tmp.Next = pre.Next
		pre.Next = tmp
	}
	return headPre.Next
}
