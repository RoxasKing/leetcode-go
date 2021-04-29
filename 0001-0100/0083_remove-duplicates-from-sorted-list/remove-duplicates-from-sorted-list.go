package main

/*
  Given the head of a sorted linked list, delete all duplicates such that each element appears only once. Return the linked list sorted as well.

  Example 1:
    Input: head = [1,1,2]
    Output: [1,2]

  Example 2:
    Input: head = [1,1,2,3,3]
    Output: [1,2,3]

  Constraints:
    1. The number of nodes in the list is in the range [0, 300].
    2. -100 <= Node.val <= 100
    3. The list is guaranteed to be sorted in ascending order.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	ptr := head
	for node := ptr.Next; node != nil; node = node.Next {
		if node.Val != ptr.Val {
			ptr.Next = node
			ptr = ptr.Next
		}
	}
	ptr.Next = nil
	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}
