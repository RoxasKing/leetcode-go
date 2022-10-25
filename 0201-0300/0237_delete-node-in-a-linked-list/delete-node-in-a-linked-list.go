package main

// Difficulty:
// Medium

// Tags:
// Linked List

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	pre := &ListNode{}
	for cur := node; cur.Next != nil; cur = cur.Next {
		cur.Val = cur.Next.Val
		pre = cur
	}
	pre.Next = nil
}
