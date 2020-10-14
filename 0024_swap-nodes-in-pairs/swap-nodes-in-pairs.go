package main

/*
  给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
  你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

  示例:
    给定 1->2->3->4, 你应该返回 2->1->4->3.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/swap-nodes-in-pairs
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func swapPairs(head *ListNode) *ListNode {
	headPre := &ListNode{Next: head}
	for n := headPre; n != nil && n.Next != nil && n.Next.Next != nil; n = n.Next.Next {
		a, b, c := n.Next, n.Next.Next, n.Next.Next.Next
		n.Next = b
		b.Next = a
		a.Next = c
	}
	return headPre.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
