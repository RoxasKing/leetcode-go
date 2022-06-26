package main

import "container/heap"

// Difficulty:
// Medium

// Tags:
// Priority Queue

func findKthLargest(nums []int, k int) int {
	h := minh{}
	for _, x := range nums {
		if h.Len() < k {
			heap.Push(&h, x)
		} else if h[0] < x {
			heap.Pop(&h)
			heap.Push(&h, x)
		}
	}
	return h[0]
}

type minh []int

func (h minh) Len() int            { return len(h) }
func (h minh) Less(i, j int) bool  { return h[i] < h[j] }
func (h minh) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minh) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *minh) Pop() interface{}   { i := h.Len() - 1; o := (*h)[i]; *h = (*h)[:i]; return o }
