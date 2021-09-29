package main

// Tags:
// Two Pointers

type ListNode struct {
	Val  int
	Next *ListNode
}

func splitListToParts(head *ListNode, k int) []*ListNode {
	n := 0
	for p := head; p != nil; p = p.Next {
		n++
	}
	a, b := n/k, n%k
	out := make([]*ListNode, k)
	for i, l, r := 0, head, head; i < k && l != nil; i++ {
		for j := 0; j < a-1; j++ {
			r = r.Next
		}
		if a > 0 && b > 0 {
			r = r.Next
			b--
		}
		next := r.Next
		r.Next = nil
		out[i] = l
		l, r = next, next
	}
	return out
}
