package main

/*
  Given the head of a linked list, rotate the list to the right by k places.

  Example 1:
    Input: head = [1,2,3,4,5], k = 2
    Output: [4,5,1,2,3]

  Example 2:
    Input: head = [0,1,2], k = 4
    Output: [2,0,1]

  Constraints:
    1. The number of nodes in the list is in the range [0, 500].
    2. -100 <= Node.val <= 100
    3. 0 <= k <= 2 * 10^9

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/rotate-list
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
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}

	cnt := 0
	for n := head; n != nil; n = n.Next {
		cnt++
	}

	k %= cnt
	if k == 0 {
		return head
	}

	last := head
	for k > 0 {
		last = last.Next
		k--
	}

	ptr := head
	for last.Next != nil {
		ptr = ptr.Next
		last = last.Next
	}

	newHead := ptr.Next
	ptr.Next = nil
	last.Next = head

	return newHead
}

type ListNode struct {
	Val  int
	Next *ListNode
}
