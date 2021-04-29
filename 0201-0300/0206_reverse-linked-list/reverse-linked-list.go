package main

/*
  Given the head of a singly linked list, reverse the list, and return the reversed list.

  Example 1:
    Input: head = [1,2,3,4,5]
    Output: [5,4,3,2,1]

  Example 2:
    Input: head = [1,2]
    Output: [2,1]

  Example 3:
    Input: head = []
    Output: []

  Constraints:
    1. The number of nodes in the list is the range [0, 5000].
    2. -5000 <= Node.val <= 5000

  Follow up: A linked list can be reversed either iteratively or recursively. Could you implement both?

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/reverse-linked-list
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
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
