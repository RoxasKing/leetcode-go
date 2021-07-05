package main

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
