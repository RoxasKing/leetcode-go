package My_LeetCode_In_Go

import (
	. "My_LeetCode_In_Go/util/linkedlist"
)

/*
  给定一个链表，旋转链表，将链表每个节点向右移动 k 个位置，其中 k 是非负数。
*/

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	var size int
	for node := head; node != nil; node = node.Next {
		size++
	}

	k = k % size
	if k == 0 {
		return head
	}

	pre := &ListNode{Val: -1, Next: head}
	ptr := pre
	for i := 0; i < size-k; i++ {
		ptr = ptr.Next
	}
	rotatingPart := ptr.Next
	ptr.Next = nil

	nonRotatingPart := pre.Next
	pre.Next = rotatingPart
	ptr = pre
	for i := 0; i < k; i++ {
		ptr = ptr.Next
	}
	ptr.Next = nonRotatingPart

	return pre.Next
}
