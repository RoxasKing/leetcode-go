package main

import "math/rand"

// Difficulty:
// Medium

// Tags:
// Reservoir Sampling

type ListNode struct {
	Val  int
	Next *ListNode
}

type Solution struct {
	head *ListNode
}

/** @param head The linked list's head.
  Note that the head is guaranteed to be not null, so it contains at least one node. */
func Constructor(head *ListNode) Solution {
	return Solution{head: head}
}

/** Returns a random node's value. */
func (this *Solution) GetRandom() int {
	var out, cnt int
	for p := this.head; p != nil; p = p.Next {
		cnt++
		if i := rand.Intn(cnt) + 1; i == cnt {
			out = p.Val
		}
	}
	return out
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
