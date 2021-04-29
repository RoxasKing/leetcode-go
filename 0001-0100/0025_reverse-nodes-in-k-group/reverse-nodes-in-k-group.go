package main

/*
  Given a linked list, reverse the nodes of a linked list k at a time and return its modified list.

  k is a positive integer and is less than or equal to the length of the linked list. If the number of nodes is not a multiple of k then left-out nodes, in the end, should remain as it is.

  Follow up:
    1. Could you solve the problem in O(1) extra memory space?
    2. You may not alter the values in the list's nodes, only nodes itself may be changed.

  Example 1:
    Input: head = [1,2,3,4,5], k = 2
    Output: [2,1,4,3,5]

  Example 2:
    Input: head = [1,2,3,4,5], k = 3
    Output: [3,2,1,4,5]

  Example 3:
    Input: head = [1,2,3,4,5], k = 1
    Output: [1,2,3,4,5]

  Example 4:
    Input: head = [1], k = 1
    Output: [1]

  Constraints:
    1. The number of nodes in the list is in the range sz.
    2. 1 <= sz <= 5000
    3. 0 <= Node.val <= 1000
    4. 1 <= k <= sz

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/reverse-nodes-in-k-group
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

func reverseKGroup(head *ListNode, k int) *ListNode {
	node := head
	nodePre := &ListNode{Next: node}
	var out *ListNode
	for {
		ptr := node
		pre := nodePre
		i := 0
		for i < k && pre.Next != nil {
			pre = pre.Next
			i++
		}
		if i == k {
			node = pre.Next
			pre.Next = nil
			first, last := revList(ptr)
			if out == nil {
				out = first
			}
			nodePre.Next = first
			last.Next = node
			nodePre = last
		} else {
			break
		}
	}
	return out
}

func revList(head *ListNode) (*ListNode, *ListNode) {
	var first, last *ListNode
	for head != nil {
		next := head.Next
		head.Next = first
		first = head
		if last == nil {
			last = head
		}
		head = next
	}
	return first, last
}

type ListNode struct {
	Val  int
	Next *ListNode
}
