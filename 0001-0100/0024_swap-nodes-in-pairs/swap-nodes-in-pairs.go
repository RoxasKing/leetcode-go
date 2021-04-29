package main

/*
  Given a linked list, swap every two adjacent nodes and return its head.

  Example 1:
    Input: head = [1,2,3,4]
    Output: [2,1,4,3]

  Example 2:
    Input: head = []
    Output: []

  Example 3:
    Input: head = [1]
    Output: [1]

  Constraints:
    1. The number of nodes in the list is in the range [0, 100].
    2. 0 <= Node.val <= 100

  Follow up: Can you solve the problem without modifying the values in the list's nodes? (i.e., Only nodes themselves may be changed.)

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/swap-nodes-in-pairs
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	headPre := &ListNode{Next: head}
	ptr := headPre
	for ptr.Next != nil && ptr.Next.Next != nil {
		next := ptr.Next.Next.Next
		a, b := ptr.Next, ptr.Next.Next
		ptr.Next = b
		b.Next = a
		a.Next = next
		ptr = a
	}
	return headPre.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
