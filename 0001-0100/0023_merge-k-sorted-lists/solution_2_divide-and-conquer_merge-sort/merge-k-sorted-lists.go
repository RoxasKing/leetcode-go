package main

// Difficulty:
// Hard

// Tags:
// Divide and conquer
// Merge Sort

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	n := len(lists)
	if n == 0 {
		return nil
	} else if n == 1 {
		return lists[0]
	}
	m := n >> 1
	return merge(mergeKLists(lists[:m]), mergeKLists(lists[m:]))
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
