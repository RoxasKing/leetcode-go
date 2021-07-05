package main

import (
	"container/heap"
	"sort"
)

// Tags:
// Priority Queue(Heap Sort)
func getSkyline(buildings [][]int) [][]int {
	n := len(buildings)
	points := make([]int, 0, n<<1)
	for _, b := range buildings {
		points = append(points, b[0], b[1])
	}
	sort.Ints(points)
	h := MaxHeap{}
	heap.Push(&h, &buildInf{rBound: 1<<63 - 1, height: 0})
	idx := 0
	out := [][]int{}

	for _, p := range points {
		for ; idx < n && buildings[idx][0] == p; idx++ {
			heap.Push(&h, &buildInf{rBound: buildings[idx][1], height: buildings[idx][2]})
		}
		for h[0].rBound <= p {
			heap.Pop(&h)
		}
		if len(out) == 0 || out[len(out)-1][1] != h[0].height {
			out = append(out, []int{p, h[0].height})
		}
	}

	return out
}

type buildInf struct {
	rBound, height int
}

type MaxHeap []*buildInf

func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Less(i, j int) bool  { return h[i].height > h[j].height }
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(*buildInf)) }
func (h *MaxHeap) Pop() interface{} {
	top := h.Len() - 1
	out := (*h)[top]
	*h = (*h)[:top]
	return out
}
