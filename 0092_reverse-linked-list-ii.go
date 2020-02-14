package leetcode

/*
  反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。
  说明:
  1 ≤ m ≤ n ≤ 链表长度。
*/

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	head = &ListNode{Val: -1, Next: head}
	leftPre := head
	left := leftPre.Next
	index := 1
	for index < m && left.Next != nil {
		leftPre = leftPre.Next
		left = left.Next
		index++
	}
	right := left
	ptr := right.Next
	index++
	for index <= n {
		tmp := ptr.Next
		leftPre.Next = ptr
		ptr.Next = left
		right.Next = tmp
		left = ptr
		ptr = tmp
		index++
	}
	return head.Next
}
