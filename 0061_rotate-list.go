package leetcode

/*
  给定一个链表，旋转链表，将链表每个节点向右移动 k 个位置，其中 k 是非负数。
*/

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	tail := head
	size := 1
	for tail.Next != nil {
		tail = tail.Next
		size++
	}
	k = k % size
	if k == 0 {
		return head
	}
	tail.Next = head // Circular-Linked-List
	for i := 1; i < size-k; i++ {
		head = head.Next
	}
	next := head.Next
	head.Next = nil
	return next
}
