package main

import "container/heap"

// Difficulty:
// Hard

// Tags:
// Priority Queue

func trapRainWater(heightMap [][]int) int {
	m, n := len(heightMap), len(heightMap[0])
	if m < 3 || n < 3 {
		return 0
	}
	vis := make([][]bool, m)
	pq := MinHeap{}
	for i, hs := range heightMap {
		vis[i] = make([]bool, n)
		for j, h := range hs {
			if i == 0 || i == m-1 || j == 0 || j == n-1 {
				vis[i][j] = true
				heap.Push(&pq, &Cell{i, j, h})
			}
		}
	}
	dirs := []int{-1, 0, 1, 0, -1}
	out := 0
	for pq.Len() > 0 {
		c := heap.Pop(&pq).(*Cell)
		for i := 0; i < 4; i++ {
			x, y := c.x+dirs[i], c.y+dirs[i+1]
			if x < 0 || m-1 < x || y < 0 || n-1 < y || vis[x][y] {
				continue
			}
			vis[x][y] = true
			h := heightMap[x][y]
			out += Max(0, c.h-h)
			heap.Push(&pq, &Cell{x, y, Max(c.h, h)})
		}
	}
	return out
}

type Cell struct{ x, y, h int }
type MinHeap []*Cell

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i].h < h[j].h }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(*Cell)) }
func (h *MinHeap) Pop() interface{} {
	i := h.Len() - 1
	out := (*h)[i]
	*h = (*h)[:i]
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
