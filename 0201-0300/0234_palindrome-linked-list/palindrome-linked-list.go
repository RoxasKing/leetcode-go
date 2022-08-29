package main

// Difficulty:
// Easy

// Tags:
// Two Pointers

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	reverse := func(head *ListNode) *ListNode {
		var o *ListNode
		for head != nil {
			next := head.Next
			head.Next = o
			o = head
			head = next
		}
		return o
	}
	p1, p2 := head, reverse(slow.Next)
	slow.Next = nil
	for ; p1 != nil && p2 != nil; p1, p2 = p1.Next, p2.Next {
		if p1.Val != p2.Val {
			return false
		}
	}
	return true
}
