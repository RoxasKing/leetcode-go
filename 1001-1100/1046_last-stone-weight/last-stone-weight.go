package main

import "container/heap"

// Difficulty:
// Easy

// Tags:
// Priority Queue

func lastStoneWeight(stones []int) int {
	stones = append(stones, 0)
	h := maxq(stones)
	heap.Init(&h)
	for h.Len() > 1 {
		a := heap.Pop(&h).(int)
		b := heap.Pop(&h).(int)
		if a != b {
			heap.Push(&h, a-b)
		}
	}
	return h[0]
}

type maxq []int

func (h maxq) Len() int            { return len(h) }
func (h maxq) Less(i, j int) bool  { return h[i] > h[j] }
func (h maxq) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *maxq) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *maxq) Pop() interface{}   { i := h.Len() - 1; o := (*h)[i]; *h = (*h)[:i]; return o }
