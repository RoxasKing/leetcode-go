package main

// Difficulty:
// Medium

// Tags:
// Two Pointers

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	headPre := &ListNode{Next: head}
	cnt, cur := 1, headPre
	for ptr := cur.Next.Next; ptr != nil; ptr = ptr.Next {
		if ptr.Val == cur.Next.Val {
			cnt++
		} else if cnt == 1 {
			cur = cur.Next
		} else {
			cur.Next = ptr
			cnt = 1
		}
	}
	if cnt > 1 {
		cur.Next = nil
	}
	return headPre.Next
}
