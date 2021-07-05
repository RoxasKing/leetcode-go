package main

import "container/heap"

// Priority Queue + Greedy Algorithm

func furthestBuilding(heights []int, bricks int, ladders int) int {
	n := len(heights)
	h := MaxHeap{}
	for i, height := range heights {
		if bricks < 0 && ladders > 0 {
			bricks += heap.Pop(&h).(int)
			ladders--
		}
		if bricks < 0 {
			return i - 1
		}
		if i+1 < n && heights[i+1] > height {
			diff := heights[i+1] - height
			bricks -= diff
			heap.Push(&h, diff)
		}
	}
	return n - 1
}

type MaxHeap []int

func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() interface{} {
	top := h.Len() - 1
	out := (*h)[top]
	*h = (*h)[:top]
	return out
}
