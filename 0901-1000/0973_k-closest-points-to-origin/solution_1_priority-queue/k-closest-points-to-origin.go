package main

import "container/heap"

// Difficulty:
// Medium

// Priority Queue

func kClosest(points [][]int, k int) [][]int {
	h := MaxHeap{}
	for _, point := range points {
		if h.Len() < k {
			heap.Push(&h, point)
		} else if distance(h[0]) > distance(point) {
			heap.Pop(&h)
			heap.Push(&h, point)
		}
	}
	return h
}

type MaxHeap [][]int

func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Less(i, j int) bool  { return distance(h[i]) > distance(h[j]) }
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.([]int)) }
func (h *MaxHeap) Pop() interface{} {
	i := h.Len() - 1
	out := (*h)[i]
	*h = (*h)[:i]
	return out
}

func distance(point []int) int {
	x, y := point[0], point[1]
	return x*x + y*y
}
