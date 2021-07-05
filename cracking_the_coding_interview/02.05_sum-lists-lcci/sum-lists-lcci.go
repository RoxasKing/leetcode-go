package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ptrPre := &ListNode{}
	ptr := ptrPre
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
		ptr.Next = &ListNode{Val: remain % 10}
		ptr = ptr.Next
		remain /= 10
	}
	if remain > 0 {
		ptr.Next = &ListNode{Val: remain}
	}
	return ptrPre.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
