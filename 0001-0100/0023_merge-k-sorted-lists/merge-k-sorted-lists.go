package main

// Tags:
// Merge Sort
// Divide and conquer
// Recursion

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	return solve(lists)
}

func solve(lists []*ListNode) *ListNode {
	n := len(lists)
	if n == 1 {
		return lists[0]
	}
	m := n >> 1
	return merge(solve(lists[:m]), solve(lists[m:]))
}

func merge(a, b *ListNode) *ListNode {
	t := &ListNode{}
	p := t
	for ; a != nil && b != nil; p = p.Next {
		if a.Val < b.Val {
			p.Next = a
			a = a.Next
		} else {
			p.Next = b
			b = b.Next
		}
	}
	if a != nil {
		p.Next = a
	} else {
		p.Next = b
	}
	return t.Next
}
