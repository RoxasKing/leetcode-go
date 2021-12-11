package main

// Difficulty:
// Easy

// Tags:
// Bit Manipulation

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	out := 0
	for ; head != nil; head = head.Next {
		out <<= 1
		out += head.Val
	}
	return out
}
