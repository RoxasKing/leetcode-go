package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	// 用于存储小于 x 的数
	min := make([]int, 0)
	// 用于存储大于等于 x 的数
	max := make([]int, 0)
	for i := head; i != nil; i = i.Next {
		if i.Val < x {
			min = append(min, i.Val)
		} else {
			max = append(max, i.Val)
		}
	}
	pre := new(ListNode)
	node := pre
	for i := range min {
		node.Next = new(ListNode)
		node = node.Next
		node.Val = min[i]
	}
	for i := range max {
		node.Next = new(ListNode)
		node = node.Next
		node.Val = max[i]
	}

	return pre.Next
}

func main() {
	head := &ListNode{Val: 1}
	next1 := &ListNode{Val: 4}
	head.Next = next1
	next2 := &ListNode{Val: 3}
	next1.Next = next2
	next3 := &ListNode{Val: 2}
	next2.Next = next3
	next4 := &ListNode{Val: 5}
	next3.Next = next4
	next5 := &ListNode{Val: 2}
	next4.Next = next5

	x := 3

	for i := head; i != nil; i = i.Next {
		fmt.Print(i.Val, ",")
	}
	fmt.Println()
	out := partition(head, x)
	for i := out; i != nil; i = i.Next {
		fmt.Print(i.Val, ",")
	}
}
