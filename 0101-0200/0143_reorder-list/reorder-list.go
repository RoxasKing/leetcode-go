package main

/*
  You are given the head of a singly linked-list. The list can be represented as:
    L0 → L1 → … → Ln - 1 → Ln
  Reorder the list to be on the following form:
    L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
  You may not modify the values in the list's nodes. Only nodes themselves may be changed.

  Example 1:
    Input: head = [1,2,3,4]
    Output: [1,4,2,3]

  Example 2:
    Input: head = [1,2,3,4,5]
    Output: [1,5,2,4,3]

  Constraints:
    1. The number of nodes in the list is in the range [1, 5 * 10^4].
    2. 1 <= Node.val <= 1000

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/reorder-list
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
func reorderList(head *ListNode) {
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	l1, l2 := head, slow.Next
	slow.Next = nil
	l2 = reverse(l2)
	for l1 != nil && l2 != nil {
		next1 := l1.Next
		next2 := l2.Next
		l1.Next = l2
		l2.Next = next1
		l1, l2 = next1, next2
	}
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
