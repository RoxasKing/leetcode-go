package main

// Difficulty:
// Easy

// Tags:
// Two Pointers

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	var out, ptr *ListNode
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			if out == nil {
				out, ptr = list1, list1
			} else {
				ptr.Next = list1
				ptr = list1
			}
			list1 = list1.Next
		} else {
			if out == nil {
				out, ptr = list2, list2
			} else {
				ptr.Next = list2
				ptr = list2
			}
			list2 = list2.Next
		}
	}
	if list1 != nil {
		ptr.Next = list1
	}
	if list2 != nil {
		ptr.Next = list2
	}
	return out
}
