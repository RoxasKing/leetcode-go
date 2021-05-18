package main

/*
  Implement an algorithm to find the kth to last element of a singly linked list. Return the value of the element.

  Note: This problem is slightly different from the original one in the book.

  Example:
    Input:  1->2->3->4->5 和 k = 2
    Output:  4

  Note:
    k is always valid.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/kth-node-from-end-of-list-lcci
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// DFS

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func kthToLast(head *ListNode, k int) int {
	x, node := dfs(head, k)
	if x != k {
		return -1
	}
	return node.Val
}

func dfs(head *ListNode, k int) (int, *ListNode) {
	if head == nil {
		return 0, head
	}
	x, node := dfs(head.Next, k)
	if x == k {
		return x, node
	}
	return x + 1, head
}

type ListNode struct {
	Val  int
	Next *ListNode
}
