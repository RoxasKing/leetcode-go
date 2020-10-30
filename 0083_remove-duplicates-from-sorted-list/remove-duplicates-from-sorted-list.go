package main

/*
  给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。

  示例 1:
    输入: 1->1->2
    输出: 1->2

  示例 2:
    输入: 1->1->2->3->3
    输出: 1->2->3

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	ptr := head
	preVal := head.Val
	for n := head.Next; n != nil; n = n.Next {
		if n.Val != preVal {
			ptr.Next = n
			ptr = n
		}
		preVal = n.Val
	}
	ptr.Next = nil
	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}
