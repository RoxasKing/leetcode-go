package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1 = reverse(l1)
	l2 = reverse(l2)
	var out *ListNode
	var remain int
	for l1 != nil || l2 != nil {
		if l1 != nil {
			remain += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			remain += l2.Val
			l2 = l2.Next
		}
		out = &ListNode{Val: remain % 10, Next: out}
		remain /= 10
	}
	if remain > 0 {
		out = &ListNode{Val: remain, Next: out}
	}
	return out
}

func reverse(head *ListNode) *ListNode {
	var out *ListNode
	for head != nil {
		next := head.Next
		head.Next = out
		out = head
		head = next
	}
	return out
}

type ListNode struct {
	Val  int
	Next *ListNode
}
