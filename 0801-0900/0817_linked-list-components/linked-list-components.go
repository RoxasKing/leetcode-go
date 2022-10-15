package main

// Difficulty:
// Medium

// Tags:
// Linked List

type ListNode struct {
	Val  int
	Next *ListNode
}

func numComponents(head *ListNode, nums []int) int {
	h := [1e4 + 1]bool{}
	for _, x := range nums {
		h[x] = true
	}
	o := 0
	for p := head; p != nil; {
		if !h[p.Val] {
			p = p.Next
			continue
		}
		for ; p != nil && h[p.Val]; p = p.Next {
		}
		o++
	}
	return o
}
