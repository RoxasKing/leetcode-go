package main

import "container/heap"

// Tags:
// Greedy
// Priority Queue

func magicTower(nums []int) int {
	inc, dec, out := 1, 0, 0
	h := MinHeap{}
	for _, num := range nums {
		inc += num
		heap.Push(&h, num)
		for inc <= 0 && h.Len() > 0 {
			x := heap.Pop(&h).(int)
			inc -= x
			dec += x
			out++
		}
	}
	if inc+dec <= 0 {
		return -1
	}
	return out
}

type MinHeap []int

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() interface{} {
	i := h.Len() - 1
	out := (*h)[i]
	*h = (*h)[:i]
	return out
}
