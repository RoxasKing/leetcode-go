package main

/*
  给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中 没有重复出现 的数字。
*/

func deleteDuplicates0082(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	pre := &ListNode{Val: -1, Next: head}
	for node := pre; node != nil; {
		if node.Next != nil &&
			node.Next.Next != nil &&
			node.Next.Val == node.Next.Next.Val {
			curNode := node.Next.Next.Next
			curVal := node.Next.Val
			for curNode != nil && curNode.Val == curVal {
				curNode = curNode.Next
			}
			node.Next = curNode
		} else {
			node = node.Next
		}
	}
	return pre.Next
}
