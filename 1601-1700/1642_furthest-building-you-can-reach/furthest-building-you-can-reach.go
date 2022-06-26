package main

import "container/heap"

// Difficulty:
// Medium

// Tags:
// Priority Queue
// Greedy

func furthestBuilding(heights []int, bricks int, ladders int) int {
	h := minh{}
	n := len(heights)
	for i := 1; i < n; i++ {
		if heights[i-1] < heights[i] {
			heap.Push(&h, heights[i]-heights[i-1])
			if ladders > 0 {
				ladders--
			} else if bricks >= h[0] {
				bricks -= heap.Pop(&h).(int)
			} else {
				return i - 1
			}
		}
	}
	return n - 1
}

type minh []int

func (h minh) Len() int            { return len(h) }
func (h minh) Less(i, j int) bool  { return h[i] < h[j] }
func (h minh) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minh) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *minh) Pop() interface{}   { i := h.Len() - 1; o := (*h)[i]; *h = (*h)[:i]; return o }
