package main

/*
  将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。 

  示例：
    输入：1->2->4, 1->3->4
    输出：1->1->2->3->4->4

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/merge-two-sorted-lists
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	pre := new(ListNode)
	cur := pre
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}
	return pre.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
