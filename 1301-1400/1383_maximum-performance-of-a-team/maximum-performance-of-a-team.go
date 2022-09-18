package main

import (
	"container/heap"
	"sort"
)

// Tags:
// Greedy
// Sorting
// Priority Queue
// Sliding Window

func maxPerformance(n int, speed []int, efficiency []int, k int) int {
	type pair struct{ s, e int }
	a := []pair{}
	for i := 0; i < n; i++ {
		a = append(a, pair{speed[i], efficiency[i]})
	}
	sort.Slice(a, func(i, j int) bool { return a[i].e > a[j].e })
	o, ss, h := -1<<31, 0, minh{}
	for i, p := range a {
		ss += p.s
		heap.Push(&h, p.s)
		if i > k-1 {
			ss -= heap.Pop(&h).(int)
		}
		o = max(o, ss*p.e)
	}
	return o % int(1e9+7)
}

type minh []int

func (h minh) Len() int            { return len(h) }
func (h minh) Less(i, j int) bool  { return h[i] < h[j] }
func (h minh) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minh) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *minh) Pop() interface{}   { i := h.Len() - 1; o := (*h)[i]; *h = (*h)[:i]; return o }

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
