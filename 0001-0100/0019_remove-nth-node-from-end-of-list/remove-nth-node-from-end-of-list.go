package main

/*
  Given the head of a linked list, remove the nth node from the end of the list and return its head.

  Follow up: Could you do this in one pass?

  Example 1:
    Input: head = [1,2,3,4,5], n = 2
    Output: [1,2,3,5]

  Example 2:
    Input: head = [1], n = 1
    Output: []

  Example 3:
    Input: head = [1,2], n = 1
    Output: [1]

  Constraints:
    1. The number of nodes in the list is sz.
    2. 1 <= sz <= 30
    3. 0 <= Node.val <= 100
    4. 1 <= n <= sz

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	headPre := &ListNode{Next: head}
	last := headPre
	for n > 0 {
		last = last.Next
		n--
	}
	ptr := headPre
	for last.Next != nil {
		last = last.Next
		ptr = ptr.Next
	}
	ptr.Next = ptr.Next.Next
	return headPre.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
