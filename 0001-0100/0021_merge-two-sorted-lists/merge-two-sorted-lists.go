package main

func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	pre := new(ListNode)
	cur := pre
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}
	return pre.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
