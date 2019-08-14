package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	tail := head
	size := 1
	// 计算链表长度, 以及获取链表尾
	for tail.Next != nil {
		size++
		tail = tail.Next
	}
	if size == 1 || k == 0 {
		return head
	}
	var max int
	if size-k == 0 {
		return head
	} else if size-k > 0 {
		max = size - k
	} else {
		max = size - k%size
	}
	for i := 0; i < max; i++ {
		tmp := head
		head = head.Next
		tail.Next = tmp
		tmp.Next = nil
		tail = tail.Next
	}

	return head
}

func main() {
	head := &ListNode{Val: 1}
	next1 := &ListNode{Val: 2}
	head.Next = next1
	next2 := &ListNode{Val: 3}
	next1.Next = next2
	next3 := &ListNode{Val: 4}
	next2.Next = next3
	next4 := &ListNode{Val: 5}
	next3.Next = next4
	for node := head; node != nil; node = node.Next {
		fmt.Print(node.Val, ".")
	}
	fmt.Println()
	head = rotateRight(head, 0)
	for node := head; node != nil; node = node.Next {
		fmt.Print(node.Val, ".")
	}
	head = nil
	head = rotateRight(head, 2)
	for node := head; node.Next != nil; node = node.Next {
		fmt.Print(node.Val, ".")
	}
}
