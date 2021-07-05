package main

func deleteNode(node *ListNode) {
	cur := node
	pre := &ListNode{Next: cur}
	for cur.Next != nil {
		cur.Val = cur.Next.Val
		cur, pre = cur.Next, pre.Next
	}
	pre.Next = nil
}

type ListNode struct {
	Val  int
	Next *ListNode
}
