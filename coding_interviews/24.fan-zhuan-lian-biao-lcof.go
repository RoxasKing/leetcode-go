package codinginterviews

/*
  定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点。

  限制：
  0 <= 节点个数 <= 5000
*/

func reverseList(head *ListNode) *ListNode {
	var ptr *ListNode
	for head != nil {
		next := head.Next
		head.Next = ptr
		ptr = head
		head = next
	}
	return ptr
}
