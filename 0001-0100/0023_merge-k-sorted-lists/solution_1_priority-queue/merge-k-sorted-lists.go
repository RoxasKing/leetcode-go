package main

import "container/heap"

// Difficulty:
// Hard

// Tags:
// Priority Queue

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	h := MinHeap{}
	for _, node := range lists {
		if node != nil {
			heap.Push(&h, node)
		}
	}
	outPre := &ListNode{}
	for ptr := outPre; h.Len() > 0; ptr = ptr.Next {
		node := heap.Pop(&h).(*ListNode)
		ptr.Next = node
		if node.Next != nil {
			heap.Push(&h, node.Next)
		}
	}
	return outPre.Next
}

type MinHeap []*ListNode

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i].Val < h[j].Val }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(*ListNode)) }
func (h *MinHeap) Pop() interface{} {
	i := h.Len() - 1
	out := (*h)[i]
	*h = (*h)[:i]
	return out
}
