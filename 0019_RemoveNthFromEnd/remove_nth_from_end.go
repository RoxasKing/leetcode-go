package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if n == 1 && head.Next == nil {
		return nil
	}

	first := &ListNode{}
	first.Next = head

	for node := first; node.Next != nil; node = node.Next {
		pre := node
		aft := node.Next.Next
		i := 0
		tmp := node.Next
		for ; i < n; i++ {
			if tmp != nil {
				tmp = tmp.Next
			}
		}
		if i == n && tmp == nil {
			pre.Next = aft
			return first.Next
		}

	}
	return nil
}

func main() {
	head := &ListNode{Val: 1}
	nex1 := &ListNode{Val: 2}
	nex2 := &ListNode{Val: 3}
	nex3 := &ListNode{Val: 4}
	nex4 := &ListNode{Val: 5}
	head.Next = nex1
	nex1.Next = nex2
	nex2.Next = nex3
	nex3.Next = nex4
	for node := head; node != nil; node = node.Next {
		fmt.Println(node.Val)
	}
	head = removeNthFromEnd(head, 1)
	fmt.Println("after:")
	for node := head; node != nil; node = node.Next {
		fmt.Println(node.Val)
	}
}
