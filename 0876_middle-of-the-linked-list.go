package leetcode

/*
  给定一个带有头结点 head 的非空单链表，返回链表的中间结点。
  如果有两个中间结点，则返回第二个中间结点。
*/

func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	return slow
}