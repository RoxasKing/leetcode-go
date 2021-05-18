package main

/*
  You have two numbers represented by a linked list, where each node contains a single digit. The digits are stored in reverse order, such that the 1's digit is at the head of the list. Write a function that adds the two numbers and returns the sum as a linked list.

  Example:
    Input: (7 -> 1 -> 6) + (5 -> 9 -> 2). That is, 617 + 295.
    Output: 2 -> 1 -> 9. That is, 912.
    Follow Up: Suppose the digits are stored in forward order. Repeat the above problem.

  Example:
    Input: (6 -> 1 -> 7) + (2 -> 9 -> 5). That is, 617 + 295.
    Output: 9 -> 1 -> 2. That is, 912.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/sum-lists-lcci
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ptrPre := &ListNode{}
	ptr := ptrPre
	remain := 0
	for l1 != nil || l2 != nil {
		if l1 != nil {
			remain += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			remain += l2.Val
			l2 = l2.Next
		}
		ptr.Next = &ListNode{Val: remain % 10}
		ptr = ptr.Next
		remain /= 10
	}
	if remain > 0 {
		ptr.Next = &ListNode{Val: remain}
	}
	return ptrPre.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
