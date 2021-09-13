package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	var out, l, r, p *ListNode
	t := &ListNode{}
	i := 0
	for p = head; p != nil; {
		next := p.Next
		p.Next = l
		l = p
		if r == nil {
			r = p
		}
		p = next
		i++
		if i == k {
			i = 0
			t.Next = l
			t = r
			if out == nil {
				out = l
			}
			l = nil
			r = nil
		}
	}
	for p = nil; l != nil; {
		next := l.Next
		l.Next = p
		p = l
		l = next
	}
	t.Next = p
	return out
}
