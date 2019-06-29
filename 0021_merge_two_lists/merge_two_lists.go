package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	head := &ListNode{}
	pre := head
	out := head
	for {
		if l1 == nil && l2 == nil {
			pre.Next = nil
			break
		} else if l1 != nil && l2 != nil {
			if l1.Val < l2.Val {
				head.Val = l1.Val
				l1 = l1.Next
			} else {
				head.Val = l2.Val
				l2 = l2.Next
			}
		} else if l1 != nil && l2 == nil {
			head.Val = l1.Val
			l1 = l1.Next
		} else if l1 == nil && l2 != nil {
			head.Val = l2.Val
			l2 = l2.Next
		}
		pre = head
		head.Next = &ListNode{}
		head = head.Next
	}
	return out
}

func main() {
	var head1 *ListNode
	//head1 := &ListNode{Val: 1}
	//nex11 := &ListNode{Val: 2}
	//nex12 := &ListNode{Val: 4}
	//head1.Next = nex11
	//nex11.Next = nex12
	//
	var head2 *ListNode
	//head2 := &ListNode{Val: 1}
	//nex21 := &ListNode{Val: 3}
	//nex22 := &ListNode{Val: 4}
	//head2.Next = nex21
	//nex21.Next = nex22

	ll := mergeTwoLists(head1, head2)

	for dd := ll; dd != nil; dd = dd.Next {
		fmt.Println(dd.Val)
	}
}
