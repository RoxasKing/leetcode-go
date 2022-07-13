package main

import "container/heap"

// Difficulty:
// Medium

// Tags:
// Priority Queue

func maxResult(nums []int, k int) int {
	n := len(nums)
	h := maxh{}
	for i, x := range nums {
		for ; h.Len() > 0 && h[0].k < i-k; heap.Pop(&h) {
		}
		v := x
		if h.Len() > 0 {
			v += h[0].v
		}
		if i == n-1 {
			return v
		}
		heap.Push(&h, pair{i, v})
	}
	return -1
}

type pair struct{ k, v int }
type maxh []pair

func (h maxh) Len() int            { return len(h) }
func (h maxh) Less(i, j int) bool  { return h[i].v > h[j].v }
func (h maxh) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *maxh) Push(x interface{}) { *h = append(*h, x.(pair)) }
func (h *maxh) Pop() interface{}   { i := h.Len() - 1; o := (*h)[i]; *h = (*h)[:i]; return o }
