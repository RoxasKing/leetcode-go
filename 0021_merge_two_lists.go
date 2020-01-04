package My_LeetCode_In_Go

/*
  将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
*/

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{-1, nil}
	node := head
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			node.Next = &ListNode{l1.Val, nil}
			l1 = l1.Next
		} else {
			node.Next = &ListNode{l2.Val, nil}
			l2 = l2.Next
		}
		node = node.Next
	}
	for l1 != nil {
		node.Next = &ListNode{l1.Val, nil}
		l1 = l1.Next
		node = node.Next
	}
	for l2 != nil {
		node.Next = &ListNode{l2.Val, nil}
		l2 = l2.Next
		node = node.Next
	}
	return head.Next
}
