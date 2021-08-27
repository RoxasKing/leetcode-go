package main

import "container/heap"

// Tags:
// Priority Queue
// Dynamic Programming

func minCost(maxTime int, edges [][]int, passingFees []int) int {
	n := len(passingFees)
	g := make([][][2]int, n)
	for _, e := range edges {
		g[e[0]] = append(g[e[0]], [2]int{e[1], e[2]})
		g[e[1]] = append(g[e[1]], [2]int{e[0], e[2]})
	}

	f := make([]int, n)
	for i := range f {
		f[i] = 1<<31 - 1
	}

	h := MinHeap{state{idx: 0, time: 0, cost: passingFees[0]}}

	for h.Len() > 0 {
		s := heap.Pop(&h).(state)
		u, time, cost := s.idx, s.time, s.cost
		if f[u] < cost {
			continue
		}
		f[u] = cost

		for _, e := range g[u] {
			v, t := e[0], e[1]
			if time+t <= maxTime {
				heap.Push(&h, state{idx: v, time: time + t, cost: cost + passingFees[v]})
			}
		}
	}

	if f[n-1] == 1<<31-1 {
		return -1
	}
	return f[n-1]
}

type state struct {
	idx, time, cost int
}

type MinHeap []state

func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool {
	if h[i].time != h[j].time {
		return h[i].time < h[j].time
	}
	return h[i].cost < h[j].cost
}
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(state)) }
func (h *MinHeap) Pop() interface{} {
	top := h.Len() - 1
	out := (*h)[top]
	*h = (*h)[:top]
	return out
}
