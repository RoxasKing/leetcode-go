package main

// Tags:
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
