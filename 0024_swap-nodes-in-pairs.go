package main

/*
  给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
  你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
*/

func swapPairs(head *ListNode) *ListNode {
	pre := &ListNode{}
	pre.Next = head
	tmp := pre
	for node := head; node != nil; {
		if node.Next != nil {
			tail := node.Next.Next
			tmp.Next = node.Next
			tmp.Next.Next = node
			tmp.Next.Next.Next = tail
			tmp = tmp.Next.Next
			node = tail
		} else {
			break
		}
	}
	return pre.Next
}
