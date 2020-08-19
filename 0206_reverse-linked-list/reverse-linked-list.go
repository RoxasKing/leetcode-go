package main

/*
  反转一个单链表。

  进阶:
    你可以迭代或递归地反转链表。你能否用两种方法解决这道题？
*/

// Iteration
func reverseList(head *ListNode) *ListNode {
	var revHead *ListNode
	for head != nil {
		next := head.Next
		head.Next = revHead
		revHead = head
		head = next
	}
	return revHead
}

// Recursion
func reverseList2(head *ListNode) *ListNode {
	res, _ := recur(head)
	return res
}

func recur(node *ListNode) (*ListNode, *ListNode) {
	if node == nil || node.Next == nil {
		return node, node
	}
	head, last := recur(node.Next)
	last.Next, node.Next = node, nil
	return head, node
}

// Recursion
func reverseList3(head *ListNode) *ListNode {
	var recur func(*ListNode) *ListNode
	recur = func(node *ListNode) *ListNode {
		if node == nil || node.Next == nil {
			return node
		}
		p := recur(node.Next)
		node.Next.Next = node
		node.Next = nil
		return p
	}
	return recur(head)
}

type ListNode struct {
	Val  int
	Next *ListNode
}
