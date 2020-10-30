package main

/*
  给定一个单链表 L：L0→L1→…→Ln-1→Ln ，
  将其重新排列后变为： L0→Ln→L1→Ln-1→L2→Ln-2→…

  你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

  示例 1:
    给定链表 1->2->3->4, 重新排列为 1->4->2->3.

  示例 2:
    给定链表 1->2->3->4->5, 重新排列为 1->5->2->4->3.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/reorder-list
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	l1, l2 := head, reverse(slow.Next)
	slow.Next = nil
	combine(l1, l2)
}

func reverse(head *ListNode) *ListNode {
	var out *ListNode
	for n := head; n != nil; {
		next := n.Next
		n.Next = out
		out = n
		n = next
	}
	return out
}

func combine(l1, l2 *ListNode) {
	for l1 != nil && l2 != nil {
		l1Next := l1.Next
		l1.Next = l2
		l2 = l2.Next
		l1.Next.Next = l1Next
		l1 = l1Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}
