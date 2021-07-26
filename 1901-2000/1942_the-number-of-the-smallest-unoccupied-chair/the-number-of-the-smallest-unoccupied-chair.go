package main

import (
	"container/heap"
	"sort"
)

// Tags:
// Priority Queue

func smallestChair(times [][]int, targetFriend int) int {
	n := len(times)
	uh := make(UnoccupiedHeap, n)
	for i := range uh {
		uh[i] = i
	}
	heap.Init(&uh)
	ts := make([][]int, n)
	for i := range times {
		ts[i] = []int{i, times[i][0], times[i][1]}
	}
	sort.Slice(ts, func(i, j int) bool { return ts[i][1] < ts[j][1] })

	oh := OccupiedHeap{}
	for t, i := 1, 0; ; t++ {
		for oh.Len() > 0 && oh[0].end == t {
			o := heap.Pop(&oh).(occupied)
			heap.Push(&uh, o.idx)
		}
		for ; i < n && ts[i][1] == t; i++ {
			idx := heap.Pop(&uh).(int)
			if ts[i][0] == targetFriend {
				return idx
			}
			heap.Push(&oh, occupied{idx: idx, end: ts[i][2]})
		}
	}
}

type UnoccupiedHeap []int

func (h UnoccupiedHeap) Len() int            { return len(h) }
func (h UnoccupiedHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h UnoccupiedHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *UnoccupiedHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *UnoccupiedHeap) Pop() interface{} {
	top := h.Len() - 1
	out := (*h)[top]
	*h = (*h)[:top]
	return out
}

type occupied struct {
	idx, end int
}

type OccupiedHeap []occupied

func (h OccupiedHeap) Len() int            { return len(h) }
func (h OccupiedHeap) Less(i, j int) bool  { return h[i].end < h[j].end }
func (h OccupiedHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *OccupiedHeap) Push(x interface{}) { *h = append(*h, x.(occupied)) }
func (h *OccupiedHeap) Pop() interface{} {
	top := h.Len() - 1
	out := (*h)[top]
	*h = (*h)[:top]
	return out
}
