package main

/*
  Write code to remove duplicates from an unsorted linked list.

  Example1:
    Input: [1, 2, 3, 3, 2, 1]
    Output: [1, 2, 3]

  Example2:
    Input: [1, 1, 1, 1, 2]
    Output: [1, 2]

  Note:
    1. The length of the list is within the range[0, 20000].
    2. The values of the list elements are within the range [0, 20000].

  Follow Up:
    How would you solve this problem if a temporary buffer is not allowed?

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/remove-duplicate-node-lcci
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeDuplicateNodes(head *ListNode) *ListNode {
	headPre := &ListNode{Next: head}
	ptr := headPre
	mark := make([]bool, 20001)
	for node := head; node != nil; node = node.Next {
		if !mark[node.Val] {
			mark[node.Val] = true
			ptr.Next = node
			ptr = node
		}
	}
	ptr.Next = nil
	return headPre.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
