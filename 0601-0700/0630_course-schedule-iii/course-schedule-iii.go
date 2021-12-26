package main

import (
	"container/heap"
	"sort"
)

// Difficulty:
// Hard

// Tags:
// Priority Queue
// Greedy

func scheduleCourse(courses [][]int) int {
	sort.Slice(courses, func(i, j int) bool { return courses[i][1] < courses[j][1] })
	h := MaxHeap{}
	current := 0
	for _, course := range courses {
		duration, lastDay := course[0], course[1]
		if current+duration <= lastDay {
			current += duration
			heap.Push(&h, duration)
		} else if h.IntSlice.Len() > 0 && duration < h.IntSlice[0] {
			current -= h.IntSlice[0] - duration
			h.IntSlice[0] = duration
			heap.Fix(&h, 0)
		}
	}
	return h.Len()
}

type MaxHeap struct{ sort.IntSlice }

func (h MaxHeap) Less(i, j int) bool  { return h.IntSlice[i] > h.IntSlice[j] }
func (h *MaxHeap) Push(x interface{}) { h.IntSlice = append(h.IntSlice, x.(int)) }
func (h *MaxHeap) Pop() interface{} {
	top := h.IntSlice.Len() - 1
	out := h.IntSlice[top]
	h.IntSlice = h.IntSlice[:top]
	return out
}
