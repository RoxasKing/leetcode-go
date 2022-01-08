package main

import (
	"container/heap"
	"sort"
)

// Difficulty:
// Medium

// Tags:
// Sorting
// Priority Queue

func carPooling(trips [][]int, capacity int) bool {
	sort.Slice(trips, func(i, j int) bool { return trips[i][1] < trips[j][1] })
	h := MinHeap{}
	for i := 0; i <= 1000 && (len(trips) > 0 || h.Len() > 0); i++ {
		for h.Len() > 0 && h[0][2] == i {
			trip := heap.Pop(&h).([]int)
			capacity += trip[0]
		}
		for ; len(trips) > 0 && trips[0][1] == i; trips = trips[1:] {
			if capacity -= trips[0][0]; capacity < 0 {
				return false
			}
			heap.Push(&h, trips[0])
		}
	}
	return h.Len() == 0
}

type MinHeap [][]int // [numPassengers, from, to]

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i][2] < h[j][2] }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.([]int)) }
func (h *MinHeap) Pop() interface{} {
	i := h.Len() - 1
	out := (*h)[i]
	*h = (*h)[:i]
	return out
}
