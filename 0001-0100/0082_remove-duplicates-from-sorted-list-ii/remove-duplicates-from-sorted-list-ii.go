package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	headPre := &ListNode{Next: head}
	ptr := headPre
	cnt := 1
	for node := ptr.Next.Next; node != nil; node = node.Next {
		if node.Val == ptr.Next.Val {
			cnt++
		} else if cnt == 1 {
			ptr = ptr.Next
		} else {
			ptr.Next = node
			cnt = 1
		}
	}
	if cnt > 1 {
		ptr.Next = nil
	}
	return headPre.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
