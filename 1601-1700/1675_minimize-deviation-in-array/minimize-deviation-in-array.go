package main

import "container/heap"

// Difficulty:
// Hard

// Tags:
// Priority Queue

func minimumDeviation(nums []int) int {
	h := make(MaxHeap, 0, len(nums))
	min := 1<<31 - 1
	for _, num := range nums {
		if num&1 == 1 {
			num <<= 1
		}
		heap.Push(&h, num)
		min = Min(min, num)
	}
	out := 1<<31 - 1
	for h.Len() > 0 {
		num := heap.Pop(&h).(int)
		out = Min(out, num-min)
		if num&1 == 1 {
			break
		}
		num >>= 1
		heap.Push(&h, num)
		min = Min(min, num)
	}
	return out
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type MaxHeap []int

func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() interface{} {
	i := h.Len() - 1
	out := (*h)[i]
	*h = (*h)[:i]
	return out
}
