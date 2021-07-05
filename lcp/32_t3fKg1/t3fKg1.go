package main

import (
	"container/heap"
	"sort"
)

// Tags:
// Priority Queue(Heap Sort)
func processTasks(tasks [][]int) int {
	sort.Slice(tasks, func(i, j int) bool { return tasks[i][0] < tasks[j][0] })
	tasks = append(tasks, []int{1e9 + 1, 1e9 + 1, 1})
	h := MinHeap{}
	out := 0
	for _, t := range tasks {
		start, end, period := t[0], t[1], t[2]
		for h.Len() > 0 && h[0][0]+out < start {
			if h[0][0]+out >= h[0][1] {
				heap.Pop(&h)
			} else {
				out += Min(h[0][1], start) - (h[0][0] + out)
			}
		}
		heap.Push(&h, [2]int{end - period + 1 - out, end + 1})
	}
	return out
}

type MinHeap [][2]int

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i][0] < h[j][0] }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.([2]int)) }
func (h *MinHeap) Pop() interface{} {
	top := h.Len() - 1
	out := (*h)[top]
	*h = (*h)[:top]
	return out
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
