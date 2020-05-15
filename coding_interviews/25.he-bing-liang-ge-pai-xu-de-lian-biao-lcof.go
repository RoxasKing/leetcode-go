package codinginterviews

/*
  输入两个递增排序的链表，合并这两个链表并使新链表中的节点仍然是递增排序的。

  限制：
  0 <= 链表长度 <= 1000
*/

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	pre := &ListNode{Val: -1}
	cur := pre
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			cur.Next = l2
			l2 = l2.Next
		} else {
			cur.Next = l1
			l1 = l1.Next
		}
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}
	return pre.Next
}
