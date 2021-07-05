package main

import "container/heap"

// Priority Queue(Heap Sort)
func magicTower(nums []int) int {
	h := MinHeap{}
	blood := 1
	times := 0
	dec := 0
	for _, num := range nums {
		heap.Push(&h, num)
		blood += num
		for h.Len() > 0 && blood <= 0 {
			v := heap.Pop(&h).(int)
			blood -= v
			dec += v
			times++
		}
	}

	if blood+dec <= 0 {
		return -1
	}
	return times
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
