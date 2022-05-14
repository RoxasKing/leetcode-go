package main

import "container/heap"

// Difficulty:
// Medium

// Tags:
// Graph
// Dijkstra
// Priority Queue

func networkDelayTime(times [][]int, n int, k int) int {
	g := make([][][2]int, n+1)
	for _, t := range times {
		u, v, w := t[0], t[1], t[2]
		g[u] = append(g[u], [2]int{v, w})
	}
	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		f[i] = 1<<31 - 1
	}
	f[k] = 0
	for h := (minh{k}); h.Len() > 0; {
		u := heap.Pop(&h).(int)
		for _, e := range g[u] {
			v, w := e[0], e[1]
			if f[u]+w < f[v] {
				f[v] = f[u] + w
				heap.Push(&h, v)
			}
		}
	}
	o := 0
	for i := 1; i <= n; i++ {
		if f[i] == 1<<31-1 {
			return -1
		}
		if o < f[i] {
			o = f[i]
		}
	}
	return o
}

type minh []int

func (h minh) Len() int            { return len(h) }
func (h minh) Less(i, j int) bool  { return h[i] < h[j] }
func (h minh) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minh) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *minh) Pop() interface{}   { i := h.Len() - 1; o := (*h)[i]; *h = (*h)[:i]; return o }
