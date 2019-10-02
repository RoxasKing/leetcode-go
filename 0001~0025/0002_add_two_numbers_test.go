package _001_0025

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	resPre := new(ListNode)
	cur := resPre
	carry := 0

	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		carry = sum / 10

		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next
	}
	return resPre.Next
}

func TestAddTwoNumbers(t *testing.T) {
	head1 := &ListNode{Val: 1}
	node1 := &ListNode{Val: 2}
	node2 := &ListNode{Val: 3}
	head1.Next = node1
	node1.Next = node2
	head2 := &ListNode{Val: 3}
	node3 := &ListNode{Val: 4}
	node4 := &ListNode{Val: 5}
	head2.Next = node3
	node3.Next = node4
	node := addTwoNumbers(head1, head2)
	for n := node; n != nil; n = n.Next {
		fmt.Print(n.Val, " ")
	}
	fmt.Println()
}
