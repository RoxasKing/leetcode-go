package main

import "math/rand"

// Reservoir Sampling

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
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
	out := 0
	cnt := 0
	for ptr := this.head; ptr != nil; ptr = ptr.Next {
		cnt++
		i := rand.Intn(cnt) + 1
		if i == cnt {
			out = ptr.Val
		}
	}
	return out
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */

type ListNode struct {
	Val  int
	Next *ListNode
}
