package leetcode

/*
  反转一个单链表。
  示例:
    输入: 1->2->3->4->5->NULL
    输出: 5->4->3->2->1->NULL
  进阶:
  你可以迭代或递归地反转链表。你能否用两种方法解决这道题？
*/

func reverseList(head *ListNode) *ListNode {
	var rev *ListNode
	for head != nil {
		next := head.Next
		head.Next = rev
		rev = head
		head = next
	}
	return rev
}
