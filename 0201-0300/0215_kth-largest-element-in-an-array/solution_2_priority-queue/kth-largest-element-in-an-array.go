package main

import "container/heap"

// Tags:
// Priority Queue

func findKthLargest(nums []int, k int) int {
	h := make(MinHeap, 0, k)
	for _, num := range nums {
		if h.Len() < k {
			heap.Push(&h, num)
		} else if h.Len() == k && h[0] < num {
			heap.Pop(&h)
			heap.Push(&h, num)
		}
	}

	return h[0]
}

type MinHeap []int

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() interface{} {
	top := h.Len() - 1
	out := (*h)[top]
	*h = (*h)[:top]
	return out
}
