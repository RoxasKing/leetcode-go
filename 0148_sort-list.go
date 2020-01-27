package My_LeetCode_In_Go

import (
	. "My_LeetCode_In_Go/util/linkedlist"
)

/*
  在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序。
*/

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	mid := slow.Next
	slow.Next = nil
	return sortList_merge(sortList(head), sortList(mid))
}

func sortList_merge(l1, l2 *ListNode) *ListNode {
	head := &ListNode{}
	cur := head
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			cur.Next, l1 = l1, l1.Next
		} else {
			cur.Next, l2 = l2, l2.Next
		}
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}
	return head.Next
}
