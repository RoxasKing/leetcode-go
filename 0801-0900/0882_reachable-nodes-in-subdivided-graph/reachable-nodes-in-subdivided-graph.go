package main

import "container/heap"

// Difficulty:
// Hard

// Tags:
// Graph
// Dijkstra
// Priority Queue
// Dynamic Programming

func reachableNodes(edges [][]int, maxMoves int, n int) int {
	g := make([][][]int, n)
	for _, e := range edges {
		u, v, cnt := e[0], e[1], e[2]+1
		g[u] = append(g[u], []int{v, cnt})
		g[v] = append(g[v], []int{u, cnt})
	}
	p := make([]int, n)
	f := make([]int, n)
	for i := 1; i < n; i++ {
		f[i] = 1<<31 - 1
	}
	h := MinHeap{&point{0, 0}}
	for h.Len() > 0 {
		x := heap.Pop(&h).(*point)
		u, c := x.u, x.c
		for _, e := range g[u] {
			v, cnt := e[0], e[1]
			if c+cnt < f[v] {
				p[v] = u
				f[v] = c + cnt
				heap.Push(&h, &point{u: v, c: f[v]})
			}
		}
	}
	out := 1
	for _, e := range edges {
		u, v, cnt := e[0], e[1], e[2]
		if f[u] > maxMoves && f[v] > maxMoves {
			continue
		}
		if p[u] == v {
			out += Min(cnt+1, maxMoves-f[v])
		} else if p[v] == u {
			out += Min(cnt+1, maxMoves-f[u])
		} else {
			out += Min(cnt, Max(0, Max(0, maxMoves-f[u])+Max(0, maxMoves-f[v])))
		}
	}
	return out
}

type point struct{ u, c int }
type MinHeap []*point

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i].c < h[j].c }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(*point)) }
func (h *MinHeap) Pop() interface{} {
	i := h.Len() - 1
	out := (*h)[i]
	*h = (*h)[:i]
	return out
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
