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
	curDay := 0
	h := maxh{}
	for _, course := range courses {
		duration, lastDay := course[0], course[1]
		if curDay+duration <= lastDay {
			curDay += duration
			heap.Push(&h, duration)
		} else if h.Len() > 0 && h[0] > duration {
			curDay -= h[0] - duration
			h[0] = duration
			heap.Fix(&h, 0)
		}
	}
	return h.Len()
}

type maxh []int

func (h maxh) Len() int            { return len(h) }
func (h maxh) Less(i, j int) bool  { return h[i] > h[j] }
func (h maxh) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *maxh) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *maxh) Pop() interface{}   { i := h.Len() - 1; o := (*h)[i]; *h = (*h)[:i]; return o }
