package main

import "container/heap"

// Difficulty:
// Hard

// Tags:
// Priority Queue

func minRefuelStops(target int, startFuel int, stations [][]int) int {
	fuel, prev, h := startFuel, 0, maxh{}
	o := 0
	for _, e := range stations {
		for fuel -= e[0] - prev; fuel < 0 && h.Len() > 0; fuel += heap.Pop(&h).(int) {
			o++
		}
		if fuel < 0 {
			return -1
		}
		heap.Push(&h, e[1])
		prev = e[0]
	}
	for fuel -= target - prev; fuel < 0 && h.Len() > 0; fuel += heap.Pop(&h).(int) {
		o++
	}
	if fuel < 0 {
		return -1
	}
	return o
}

type maxh []int

func (h maxh) Len() int            { return len(h) }
func (h maxh) Less(i, j int) bool  { return h[i] > h[j] }
func (h maxh) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *maxh) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *maxh) Pop() interface{}   { i := h.Len() - 1; o := (*h)[i]; *h = (*h)[:i]; return o }
