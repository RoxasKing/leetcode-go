package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reversePrint(head *ListNode) []int {
	out := []int{}
	for p := head; p != nil; p = p.Next {
		out = append(out, p.Val)
	}
	for i := 0; i < len(out)>>1; i++ {
		out[i], out[len(out)-1-i] = out[len(out)-1-i], out[i]
	}
	return out
}

type ListNode struct {
	Val  int
	Next *ListNode
}
