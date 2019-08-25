package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	// 记录重复节点的前一个节点
	pre := &ListNode{}
	out := pre
	pre.Next = head
	// 记录上下文相同节点数量
	count := 1
	node := head
	for ; node.Next != nil; node = node.Next {
		if node.Next.Val == node.Val {
			// 如果下一个节点和当前节点相同，计数器加 1
			count++
		} else {
			// 如果不同
			if count > 1 {
				// 如果计数器大于 1，则将相同节点的前一个节点和相同节点的后一个节点相连
				pre.Next = node
				// 重置计数器
				count = 1
			}
			pre = pre.Next
		}
	}
	if count > 1 {
		pre.Next = node
	}
	return out.Next
}

func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	node := head
	for node != nil && node.Next != nil {
		if node.Val == node.Next.Val {
			node.Next = node.Next.Next
		} else {
			node = node.Next
		}
	}
	return head
}

func main() {
	head := &ListNode{Val: 1}
	next1 := &ListNode{Val: 1}
	head.Next = next1
	next2 := &ListNode{Val: 1}
	next1.Next = next2
	next3 := &ListNode{Val: 2}
	next2.Next = next3
	next4 := &ListNode{Val: 3}
	next3.Next = next4
	next5 := &ListNode{Val: 4}
	next4.Next = next5
	next6 := &ListNode{Val: 5}
	next5.Next = next6
	next7 := &ListNode{Val: 5}
	next6.Next = next7
	for i := head; i != nil; i = i.Next {
		fmt.Print(i.Val, ",")
	}
	fmt.Println()
	head = deleteDuplicates(head)
	for i := head; i != nil; i = i.Next {
		fmt.Print(i.Val, ",")
	}
}
