package leetcode

/*
  给定一个链表和一个特定值 x，对链表进行分隔，使得所有小于 x 的节点都在大于或等于 x 的节点之前。
  你应当保留两个分区中每个节点的初始相对位置。
*/

func partition(head *ListNode, x int) *ListNode {
	leftHead := &ListNode{Val: -1}
	leftTail := leftHead
	rightHead := &ListNode{Val: -1}
	rightTail := rightHead
	for node := head; node != nil; node = node.Next {
		if node.Val < x {
			leftTail.Next = node
			leftTail = leftTail.Next
		} else {
			rightTail.Next = node
			rightTail = rightTail.Next
		}
	}
	leftTail.Next = rightHead.Next
	rightTail.Next = nil
	return leftHead.Next
}
