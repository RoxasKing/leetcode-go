package main

import (
	"container/heap"
	"math"
)

// Tags:
// Priority Queue(Heap Sort)
func storeWater(bucket []int, vat []int) int {
	h := MaxHeap{}
	op := 0
	for i := range bucket {
		if vat[i] == 0 {
			continue
		}
		if bucket[i] == 0 {
			bucket[i] = 1
			op++
		}
		heap.Push(&h, [2]int{bucket[i], vat[i]})
	}

	if h.Len() == 0 {
		return 0
	}

	out := 1<<31 - 1
	for {
		e := heap.Pop(&h).([2]int)
		out = Min(out, op+e[1]/e[0]+remain(e[1], e[0]))
		if e[0] < int(math.Sqrt(float64(e[1]))) {
			op++
			heap.Push(&h, [2]int{e[0] + 1, e[1]})
		} else {
			break
		}
	}
	return out
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func remain(a, b int) int {
	if a%b == 0 {
		return 0
	}
	return 1
}

type MaxHeap [][2]int

func (h MaxHeap) Len() int { return len(h) }
func (h MaxHeap) Less(i, j int) bool {
	return float64(h[i][1])/float64(h[i][0]) > float64(h[j][1])/float64(h[j][0])
}
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.([2]int)) }
func (h *MaxHeap) Pop() interface{} {
	top := h.Len() - 1
	out := (*h)[top]
	*h = (*h)[:top]
	return out
}
