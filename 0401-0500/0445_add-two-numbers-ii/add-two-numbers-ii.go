package main

/*
  You are given two non-empty linked lists representing two non-negative integers. The most significant digit comes first and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

  You may assume the two numbers do not contain any leading zero, except the number 0 itself.

  Example 1:
    Input: l1 = [7,2,4,3], l2 = [5,6,4]
    Output: [7,8,0,7]

  Example 2:
    Input: l1 = [2,4,3], l2 = [5,6,4]
    Output: [8,0,7]

  Example 3:
    Input: l1 = [0], l2 = [0]
    Output: [0]

  Constraints:
    1. The number of nodes in each linked list is in the range [1, 100].
    2. 0 <= Node.val <= 9
    3. It is guaranteed that the list represents a number that does not have leading zeros.

  Follow up: Could you solve it without reversing the input lists?

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/add-two-numbers-ii
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
	l1 = reverse(l1)
	l2 = reverse(l2)
	var out *ListNode
	var remain int
	for l1 != nil || l2 != nil {
		if l1 != nil {
			remain += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			remain += l2.Val
			l2 = l2.Next
		}
		out = &ListNode{Val: remain % 10, Next: out}
		remain /= 10
	}
	if remain > 0 {
		out = &ListNode{Val: remain, Next: out}
	}
	return out
}

func reverse(head *ListNode) *ListNode {
	var out *ListNode
	for head != nil {
		next := head.Next
		head.Next = out
		out = head
		head = next
	}
	return out
}

type ListNode struct {
	Val  int
	Next *ListNode
}
