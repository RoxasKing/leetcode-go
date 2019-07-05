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

func swapPairs(head *ListNode) *ListNode {
	var out *ListNode
	// 定义要交换节点的前一个结点
	pre := &ListNode{}
	pre.Next = head
	out = pre
	for {
		if head != nil {
			if head.Next != nil {
				// 要交换的两个节点的后一个节点
				next := head.Next.Next
				// 交换两个节点的指针操作
				pre.Next = head.Next
				pre.Next.Next = head
				head.Next = next

				// 准备下一次交换的前一个节点
				pre = pre.Next.Next
				head = pre.Next
			} else {
				return out.Next
			}
		} else {
			return out.Next
		}
	}
}

func main() {
	head1 := &ListNode{Val: 1}
	nex1 := &ListNode{Val: 2}
	head1.Next = nex1
	nex2 := &ListNode{Val: 3}
	nex1.Next = nex2
	nex3 := &ListNode{Val: 4}
	nex2.Next = nex3
	nex4 := &ListNode{Val: 5}
	nex3.Next = nex4

	new := swapPairs(head1)

	for node := new; node != nil; node = node.Next {
		fmt.Println(node.Val)
	}
}
