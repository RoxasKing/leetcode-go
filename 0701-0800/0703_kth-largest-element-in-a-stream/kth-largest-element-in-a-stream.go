package main

import "container/heap"

// Difficulty:
// Easy

// Tags:
// Priority Queue

type KthLargest struct {
	k int
	h minh
}

func Constructor(k int, nums []int) KthLargest {
	h := minh{}
	for _, x := range nums {
		if h.Len() < k {
			heap.Push(&h, x)
		} else if x > h[0] {
			heap.Pop(&h)
			heap.Push(&h, x)
		}
	}
	return KthLargest{k, h}
}

func (this *KthLargest) Add(val int) int {
	if this.h.Len() < this.k {
		heap.Push(&this.h, val)
	} else if val > this.h[0] {
		heap.Pop(&this.h)
		heap.Push(&this.h, val)
	}
	return this.h[0]
}

type minh []int

func (h minh) Len() int            { return len(h) }
func (h minh) Less(i, j int) bool  { return h[i] < h[j] }
func (h minh) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minh) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *minh) Pop() interface{}   { i := len(*h) - 1; o := (*h)[i]; *h = (*h)[:i]; return o }

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
