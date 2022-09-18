package main

import (
	"container/heap"
	"math"
	"sort"
)

// Difficulty:
// Hard

// Tags:
// Greedy
// Sorting
// Priority Queue
// Sliding Window

func mincostToHireWorkers(quality []int, wage []int, k int) float64 {
	type pair struct{ q, w int }
	a := []pair{}
	for i := range quality {
		a = append(a, pair{quality[i], wage[i]})
	}
	sort.Slice(a, func(i, j int) bool { return a[i].w*a[j].q < a[j].w*a[i].q })
	o, sumq, h := math.MaxFloat64, 0, maxh{}
	for _, p := range a {
		sumq += p.q
		heap.Push(&h, p.q)
		if h.Len() > k {
			sumq -= heap.Pop(&h).(int)
		}
		if h.Len() == k {
			o = math.Min(o, float64(sumq*p.w)/float64(p.q))
		}
	}
	return o
}

type maxh []int

func (h maxh) Len() int            { return len(h) }
func (h maxh) Less(i, j int) bool  { return h[i] > h[j] }
func (h maxh) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *maxh) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *maxh) Pop() interface{}   { i := h.Len() - 1; o := (*h)[i]; *h = (*h)[:i]; return o }
