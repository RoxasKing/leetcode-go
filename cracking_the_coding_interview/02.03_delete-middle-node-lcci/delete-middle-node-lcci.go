package main

/*
  Implement an algorithm to delete a node in the middle (i.e., any node but the first and last node, not necessarily the exact middle) of a singly linked list, given only access to that node.

  Example:
    Input: the node c from the linked list a->b->c->d->e->f
    Output: nothing is returned, but the new linked list looks like a->b->d->e->f

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/delete-middle-node-lcci
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(node *ListNode) {
	for {
		node.Val = node.Next.Val
		if node.Next.Next == nil {
			break
		}
		node = node.Next
	}
	node.Next = nil
}

type ListNode struct {
	Val  int
	Next *ListNode
}
