package main

/*
  Write code to partition a linked list around a value x, such that all nodes less than x come before all nodes greater than or equal to x. If x is contained within the list, the values of x only need to be after the elements less than x (see below). The partition element x can appear anywhere in the "right partition"; it does not need to appear between the left and right partitions.

  Example:
    Input: head = 3->5->8->5->10->2->1, x = 5
    Output: 3->1->2->10->5->5->8

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/partition-list-lcci
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	headPre := &ListNode{Next: head}
	for ptr := headPre; ptr.Next != nil; {
		if ptr.Next.Val < x && ptr != headPre {
			node := ptr.Next
			ptr.Next = ptr.Next.Next
			node.Next = headPre.Next
			headPre.Next = node
		} else {
			ptr = ptr.Next
		}
	}
	return headPre.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
