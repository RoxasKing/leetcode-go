package main

// Difficulty:
// Medium

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	headPre := &ListNode{Next: head}
	for p := headPre; p != nil && p.Next != nil && p.Next.Next != nil; {
		a, b := p.Next, p.Next.Next
		a.Next = b.Next
		p.Next = b
		b.Next = a
		p = a
	}
	return headPre.Next
}
