package main

// Difficulty:
// Medium

// Tags:
// Two Pointers

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapNodes(head *ListNode, k int) *ListNode {
	p := &ListNode{Next: head}
	p1, p2, ed := p, p, p
	for i := 0; i < k; i++ {
		p1 = p1.Next
		ed = ed.Next
	}
	for ed != nil {
		p2 = p2.Next
		ed = ed.Next
	}
	p1.Val, p2.Val = p2.Val, p1.Val
	return p.Next
}
