package main

// Difficulty:
// Medium

// Tags:
// Linked List
// Insertion Sort

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	outPre := &ListNode{}
	for cur := head; cur != nil; {
		next := cur.Next
		cur.Next = nil
		ptr := outPre
		for ptr.Next != nil && ptr.Next.Val < cur.Val {
			ptr = ptr.Next
		}
		tmp := ptr.Next
		ptr.Next = cur
		cur.Next = tmp
		cur = next
	}
	return outPre.Next
}
