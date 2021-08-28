package main

import (
	"container/heap"
	"sort"
)

// Tags:
// Greedy
// Priority Queue

func processTasks(tasks [][]int) int {
	sort.Slice(tasks, func(i, j int) bool { return tasks[i][0] < tasks[j][0] })
	tasks = append(tasks, []int{1e9 + 1, 1e9 + 1, 1})
	h := MinHeap{}
	out := 0
	for _, t := range tasks {
		start, end, period := t[0], t[1], t[2]
		for h.Len() > 0 && h[0].st+out < start {
			if h[0].st+out < h[0].ed {
				out += Min(h[0].ed, start) - (h[0].st + out)
			} else {
				heap.Pop(&h)
			}
		}
		heap.Push(&h, pair{st: end - period + 1 - out, ed: end + 1})
	}
	return out
}

type pair struct {
	st, ed int
}

type MinHeap []pair

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i].st < h[j].st }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(pair)) }
func (h *MinHeap) Pop() interface{} {
	i := h.Len() - 1
	out := (*h)[i]
	*h = (*h)[:i]
	return out
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
