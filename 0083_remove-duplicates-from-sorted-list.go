package main

/*
  给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。
*/

func deleteDuplicates0083(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	for node := head; node != nil; node = node.Next {
		curNode := node.Next
		for curNode != nil && curNode.Val == node.Val {
			curNode = curNode.Next
		}
		node.Next = curNode
	}
	return head
}
