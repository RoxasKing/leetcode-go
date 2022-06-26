package main

import "container/heap"

// Difficulty:
// Hard

// Tags:
// Priority Queue

func isPossible(target []int) bool {
	sum := 0
	h := maxh{}
	for _, x := range target {
		sum += x
		if x > 1 {
			heap.Push(&h, x)
		}
	}
	for h.Len() > 0 {
		x := heap.Pop(&h).(int)
		t := sum - x
		if t >= x || t == 0 {
			return false
		}
		if x %= t; x == 0 {
			x += t
		}
		sum = t + x
		if x > 1 {
			heap.Push(&h, x)
		}
	}
	return true
}

type maxh []int

func (h maxh) Len() int            { return len(h) }
func (h maxh) Less(i, j int) bool  { return h[i] > h[j] }
func (h maxh) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *maxh) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *maxh) Pop() interface{}   { i := h.Len() - 1; o := (*h)[i]; *h = (*h)[:i]; return o }
