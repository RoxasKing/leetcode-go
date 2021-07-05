package main

import "container/heap"

// Dynamic Programming
// Priority Queue

func nthSuperUglyNumber(n int, primes []int) int {
	h := MinHeap{}
	for _, prime := range primes {
		heap.Push(&h, &Pointer{idx: 0, mul: prime})
	}
	for i := 1; i < n; i++ {
		for f[h[0].idx]*h[0].mul <= f[i-1] {
			p := heap.Pop(&h).(*Pointer)
			heap.Push(&h, &Pointer{idx: p.idx + 1, mul: p.mul})
		}
		f[i] = f[h[0].idx] * h[0].mul
	}
	return f[n-1]
}

var f = [100000]int{1}

type Pointer struct {
	idx int
	mul int
}

type MinHeap []*Pointer

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return f[h[i].idx]*h[i].mul < f[h[j].idx]*h[j].mul }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(*Pointer)) }
func (h *MinHeap) Pop() interface{} {
	top := h.Len() - 1
	out := (*h)[top]
	*h = (*h)[:top]
	return out
}
