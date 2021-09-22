package main

// Tags:
// Two Pointers

type ListNode struct {
	Val  int
	Next *ListNode
}

func splitListToParts(head *ListNode, k int) []*ListNode {
	cnt := 0
	for p := head; p != nil; p = p.Next {
		cnt++
	}
	a, b := cnt/k, cnt%k
	tmp := &ListNode{Next: head}
	out := make([]*ListNode, 0, k)
	for l, r := tmp, tmp; l != nil && len(out) < k; {
		for i := 0; i < a && r != nil; i++ {
			r = r.Next
		}
		if b > 0 {
			r = r.Next
			b--
		}
		out = append(out, l.Next)
		l.Next = nil
		l = r
	}
	return out
}
