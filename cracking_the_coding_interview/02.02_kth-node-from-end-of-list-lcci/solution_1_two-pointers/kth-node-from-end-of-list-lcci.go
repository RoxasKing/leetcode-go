package main

/*
  Implement an algorithm to find the kth to last element of a singly linked list. Return the value of the element.

  Note: This problem is slightly different from the original one in the book.

  Example:
    Input:  1->2->3->4->5 和 k = 2
    Output:  4

  Note:
    k is always valid.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/kth-node-from-end-of-list-lcci
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Two Pointers

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func kthToLast(head *ListNode, k int) int {
	last := head
	for ; k > 1 && last != nil; k-- {
		last = last.Next
	}
	if last == nil {
		return -1
	}
	out := head
	for last.Next != nil {
		out = out.Next
		last = last.Next
	}
	return out.Val
}

type ListNode struct {
	Val  int
	Next *ListNode
}
