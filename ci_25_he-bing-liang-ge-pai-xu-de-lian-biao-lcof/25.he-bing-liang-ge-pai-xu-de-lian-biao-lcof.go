package main

/*
  输入两个递增排序的链表，合并这两个链表并使新链表中的节点仍然是递增排序的。

  示例1：
    输入：1->2->4, 1->3->4
    输出：1->1->2->3->4->4

  限制：
    0 <= 链表长度 <= 1000

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/he-bing-liang-ge-pai-xu-de-lian-biao-lcof
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	outPre := &ListNode{}
	ptr := outPre
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			ptr.Next = l1
			l1 = l1.Next
		} else {
			ptr.Next = l2
			l2 = l2.Next
		}
		ptr = ptr.Next
	}
	if l1 != nil {
		ptr.Next = l1
	} else if l2 != nil {
		ptr.Next = l2
	}
	return outPre.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
