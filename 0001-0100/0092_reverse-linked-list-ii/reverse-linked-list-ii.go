package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	headPre := &ListNode{Next: head}
	revPre := headPre
	for i := 1; i < m; i++ {
		revPre = revPre.Next
	}
	node := revPre.Next
	var revHead, revTail *ListNode
	for i := m; i <= n; i++ {
		next := node.Next
		node.Next = revHead
		revHead = node
		if revTail == nil {
			revTail = node
		}
		node = next
	}
	revPre.Next = revHead
	revTail.Next = node
	return headPre.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
