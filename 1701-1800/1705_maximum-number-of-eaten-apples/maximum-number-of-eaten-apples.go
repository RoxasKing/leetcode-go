package main

import "container/heap"

// Difficulty:
// Medium

// Tags:
// Greedy
// Priority Queue

func eatenApples(apples []int, days []int) int {
	h := MinHeap{}
	var out int
	for day := 0; ; day++ {
		for ; h.Len() > 0 && (h[0][1] < day || h[0][0] <= 0); heap.Pop(&h) {
		}
		if day < len(apples) && apples[day] != 0 {
			heap.Push(&h, []int{apples[day], day + days[day] - 1})
		}
		if h.Len() == 0 {
			if day >= len(apples) {
				break
			}
			continue
		}
		e := heap.Pop(&h).([]int)
		e[0]--
		heap.Push(&h, e)
		out++
	}
	return out
}

type MinHeap [][]int

func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool {
	return h[i][1]-h[i][0] < h[j][1]-h[j][0] || h[i][1]-h[i][0] == h[j][1]-h[j][0] && h[i][0] < h[j][0]
}
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.([]int)) }
func (h *MinHeap) Pop() interface{} {
	i := h.Len() - 1
	out := (*h)[i]
	*h = (*h)[:i]
	return out
}
