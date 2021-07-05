package main

// Tags:
// DFS + Merge Sort

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
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
	} else if l2 != nil {
		ptr.Next = l2
	}
	return outPre.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
