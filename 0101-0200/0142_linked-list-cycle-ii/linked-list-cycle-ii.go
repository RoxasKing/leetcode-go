package main

/*
  Given a linked list, return the node where the cycle begins. If there is no cycle, return null.

  There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer. Internally, pos is used to denote the index of the node that tail's next pointer is connected to. Note that pos is not passed as a parameter.

  Notice that you should not modify the linked list.

  Example 1:
    Input: head = [3,2,0,-4], pos = 1
    Output: tail connects to node index 1
    Explanation: There is a cycle in the linked list, where tail connects to the second node.

  Example 2:
    Input: head = [1,2], pos = 0
    Output: tail connects to node index 0
    Explanation: There is a cycle in the linked list, where tail connects to the first node.

  Example 3:
    Input: head = [1], pos = -1
    Output: no cycle
    Explanation: There is no cycle in the linked list.

  Constraints:
    1. The number of the nodes in the list is in the range [0, 10^4].
    2. -10^5 <= Node.val <= 10^5
    3. pos is -1 or a valid index in the linked-list.

  Follow up: Can you solve it using O(1) (i.e. constant) memory?

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/linked-list-cycle-ii
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
