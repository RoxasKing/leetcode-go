package main

// Tags:
// Merge Sort

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
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
	outPre := &ListNode{}
	ptr := outPre
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
	} else {
		ptr.Next = l2
	}
	return outPre.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
