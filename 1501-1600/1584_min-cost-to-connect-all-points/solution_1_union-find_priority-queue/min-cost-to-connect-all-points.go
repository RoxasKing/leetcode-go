package main

import "container/heap"

// Difficult:
// Medium

// Tags:
// Union-Find
// Priority Queue

func minCostConnectPoints(points [][]int) int {
	h := minh{}
	n := len(points)
	for a := 0; a < n; a++ {
		for b := a + 1; b < n; b++ {
			d := abs(points[a][0]-points[b][0]) + abs(points[a][1]-points[b][1])
			heap.Push(&h, edge{a, b, d})
		}
	}
	o := 0
	uf := newUnionFind(n)
	for c := 0; c < n-1; {
		e := heap.Pop(&h).(edge)
		a, b := e.a, e.b
		if uf.find(a) == uf.find(b) {
			continue
		}
		uf.union(a, b)
		o += e.d
		c++
	}
	return o
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type edge struct{ a, b, d int }
type minh []edge

func (h minh) Len() int            { return len(h) }
func (h minh) Less(i, j int) bool  { return h[i].d < h[j].d }
func (h minh) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minh) Push(x interface{}) { *h = append(*h, x.(edge)) }
func (h *minh) Pop() interface{}   { i := h.Len() - 1; o := (*h)[i]; *h = (*h)[:i]; return o }

type unionFind struct{ parent, size []int }

func newUnionFind(n int) unionFind {
	parent := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}
	return unionFind{parent, size}
}

func (uf unionFind) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf unionFind) union(x, y int) {
	x, y = uf.find(x), uf.find(y)
	if x == y {
		return
	}
	if uf.size[x] < uf.size[y] {
		x, y = y, x
	}
	uf.parent[y] = x
	uf.size[x] += uf.size[y]
}
