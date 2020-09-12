package main

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
	l1, l2 := head, slow.Next
	slow.Next = nil
	return merge(sortList(l1), sortList(l2))
}

func merge(l1, l2 *ListNode) *ListNode {
	rootPre := &ListNode{}
	ptr := rootPre
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			ptr.Next = l1
			l1 = l1.Next
		} else {
			ptr.Next = l2
			l2 = l2.Next
		}
		ptr = ptr.Next
	}
	if l1 != nil {
		ptr.Next = l1
	}
	if l2 != nil {
		ptr.Next = l2
	}
	return rootPre.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
