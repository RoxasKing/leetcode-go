package main

// Difficulty:
// Medium

// Tags:
// Two Pointers

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	headPre := &ListNode{-1, head}
	p2 := headPre
	for ; n > 0; n-- {
		p2 = p2.Next
	}
	p1 := headPre
	for ; p2.Next != nil; p2 = p2.Next {
		p1 = p1.Next
	}
	p1.Next = p1.Next.Next
	return headPre.Next
}
