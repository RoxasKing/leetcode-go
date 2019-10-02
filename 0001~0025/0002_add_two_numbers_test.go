package _001_0025

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// 给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，
// 并且它们的每个节点只能存储 一位 数字。
// 如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
// 您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{0, nil}
	lastNode := res
	rest := 0
	for {
		val := rest
		if l1 != nil {
			val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val += l2.Val
			l2 = l2.Next
		}
		if val == 0 && l1 == nil && l2 == nil {
			if res.Next == nil {
				return res
			}
			return res.Next
		}
		if val >= 10 {
			val -= 10
			rest = 1
		} else {
			rest = 0
		}
		partSum := ListNode{val, nil}
		lastNode.Next = &partSum
		lastNode = lastNode.Next
	}
}

func TestAddTwoNumbers(t *testing.T) {
	head1 := &ListNode{Val: 1}
	node1 := &ListNode{Val: 2}
	node2 := &ListNode{Val: 3}
	head1.Next = node1
	node1.Next = node2
	head2 := &ListNode{Val: 4}
	node3 := &ListNode{Val: 5}
	node4 := &ListNode{Val: 6}
	head2.Next = node3
	node3.Next = node4
	node := addTwoNumbers(head1, head2)
	for n := node; n != nil; n = n.Next {
		fmt.Print(n.Val, " ")
	}
	fmt.Println()
}
