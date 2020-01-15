package My_LeetCode_In_Go

/*
  给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。
  k 是一个正整数，它的值小于或等于链表的长度。
  如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
  示例 :
  给定这个链表：1->2->3->4->5
  当 k = 2 时，应当返回: 2->1->4->3->5
  当 k = 3 时，应当返回: 3->2->1->4->5
  说明 :
    你的算法只能使用常数的额外空间。
    你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
*/

func reverseKGroup(head *ListNode, k int) *ListNode {
	pre := &ListNode{}
	tail := pre
	for {
		stack := make([]*ListNode, 0, k)
		var i int
		for ; i < k && head != nil; i++ {
			stack = append(stack, head)
			head = head.Next
		}
		if i == k {
			for j := 0; j < i; j++ {
				node := stack[len(stack)-1]
				tail.Next = node
				tail = tail.Next
				stack = stack[:len(stack)-1]
			}
			tail.Next = nil
		} else {
			for _, node := range stack {
				tail.Next = node
				tail = tail.Next
			}
			return pre.Next
		}
	}
}
