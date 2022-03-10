package main

// Difficulty:
// Medium

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var out, ptr *ListNode
	remain := 0
	for l1 != nil || l2 != nil {
		if l1 != nil {
			remain += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			remain += l2.Val
			l2 = l2.Next
		}
		if ptr != nil {
			ptr.Next = &ListNode{Val: remain % 10}
			ptr = ptr.Next
		} else {
			ptr = &ListNode{Val: remain % 10}
			out = ptr
		}
		remain /= 10
	}
	if remain > 0 {
		ptr.Next = &ListNode{Val: remain}
	}
	return out
}
