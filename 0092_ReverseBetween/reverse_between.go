package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	out := &ListNode{}
	out.Next = head
	count := 0
	var pre, aft, end *ListNode
	for node := out; count <= n; node = node.Next {
		if count == m-1 {
			pre = node
		} else if count == m {
			aft, end = node, node
		} else if count > m {
			tmp := node
			end.Next = node.Next
			node = end
			pre.Next = tmp
			tmp.Next = aft
			aft = pre.Next
		}
		count++
	}
	return out.Next
}

func main() {
	head := &ListNode{Val: 1}
	nex1 := &ListNode{Val: 2}
	head.Next = nex1
	nex2 := &ListNode{Val: 3}
	nex1.Next = nex2
	nex3 := &ListNode{Val: 4}
	nex2.Next = nex3
	nex4 := &ListNode{Val: 5}
	nex3.Next = nex4
	m, n := 2, 4

	for n := head; n != nil; n = n.Next {
		fmt.Print(n.Val, ",")
	}
	fmt.Println()

	reverseBetween(head, m, n)

	for n := head; n != nil; n = n.Next {
		fmt.Print(n.Val, ",")
	}
	fmt.Println()
}
