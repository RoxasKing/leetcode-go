package main

import (
	"container/heap"
	"math"
)

// Tags:
// Greedy
// Priority Queue

func storeWater(bucket []int, vat []int) int {
	op := 0
	h := &MaxHeap{}
	for i := range bucket {
		if vat[i] == 0 {
			continue
		}
		if bucket[i] == 0 {
			bucket[i] = 1
			op++
		}
		heap.Push(h, pair{b: bucket[i], v: vat[i]})
	}

	if h.Len() == 0 {
		return 0
	}

	out := 1<<31 - 1
	for {
		p := heap.Pop(h).(pair)
		out = Min(out, op+p.v/p.b+x(p.v, p.b))
		if p.b >= int(math.Sqrt(float64(p.v))) {
			break
		}
		op++
		heap.Push(h, pair{b: p.b + 1, v: p.v})
	}
	return out
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func x(a, b int) int {
	if a%b == 0 {
		return 0
	}
	return 1
}

type pair struct {
	b, v int
}

type MaxHeap []pair

func (h MaxHeap) Len() int { return len(h) }
func (h MaxHeap) Less(i, j int) bool {
	return float64(h[i].v)/float64(h[i].b) > float64(h[j].v)/float64(h[j].b)
}
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(pair)) }
func (h *MaxHeap) Pop() interface{} {
	i := h.Len() - 1
	out := (*h)[i]
	*h = (*h)[:i]
	return out
}
