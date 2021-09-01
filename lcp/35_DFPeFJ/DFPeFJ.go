package main

import (
	"container/heap"
	"runtime/debug"
)

// Tags:
// Dijkstra
// Priority Queue

func electricCarPlan(paths [][]int, cnt int, start int, end int, charge []int) int {
	n := len(charge)
	g := make([][]*pair, n)
	for _, p := range paths {
		u, v, w := p[0], p[1], p[2]
		g[u] = append(g[u], &pair{v: v, w: w})
		g[v] = append(g[v], &pair{v: u, w: w})
	}

	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, cnt+1)
		for j := range f[i] {
			f[i][j] = 1e9
		}
	}
	f[start][0] = 0

	h := MinHeap{&state{t: 0, u: start, e: 0}}
	for h.Len() > 0 {
		x := heap.Pop(&h).(*state)
		t, u, e := x.t, x.u, x.e
		if t > f[u][e] {
			continue
		}
		if u == end {
			return t
		}
		if e < cnt {
			tt := t + charge[u]
			if tt < f[u][e+1] {
				f[u][e+1] = tt
				heap.Push(&h, &state{t: tt, u: u, e: e + 1})
			}
		}
		for _, x := range g[u] {
			v, w := x.v, x.w
			if w <= e && t+w < f[v][e-w] {
				f[v][e-w] = t + w
				heap.Push(&h, &state{t: t + w, u: v, e: e - w})
			}
		}
	}

	return -1
}

type pair struct {
	v, w int
}

type state struct {
	t, u, e int
}

type MinHeap []*state

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i].t < h[j].t }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(*state)) }
func (h *MinHeap) Pop() interface{} {
	i := h.Len() - 1
	out := (*h)[i]
	*h = (*h)[:i]
	return out
}

func init() { debug.SetGCPercent(-1) }
