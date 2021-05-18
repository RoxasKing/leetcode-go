package main

/*
  Given a circular linked list, implement an algorithm that returns the node at the beginning of the loop.

  Circular linked list: A (corrupt) linked list in which a node's next pointer points to an earlier node, so as to make a loop in the linked list.

  Example 1:
    Input: head = [3,2,0,-4], pos = 1
    Output: tail connects to node index 1

  Example 2:
    Input: head = [1,2], pos = 0
    Output: tail connects to node index 0

  Example 3:
    Input: head = [1], pos = -1
    Output: no cycle

  Follow Up:
    Can you solve it without using additional space?

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/linked-list-cycle-lcci
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
		if slow == fast {
			break
		}
	}
	if fast == nil || fast.Next == nil {
		return nil
	}
	slow = head
	for slow != fast {
		slow, fast = slow.Next, fast.Next
	}
	return slow
}

type ListNode struct {
	Val  int
	Next *ListNode
}
