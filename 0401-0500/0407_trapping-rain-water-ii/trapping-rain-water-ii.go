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
	pq := &PriorityQueue{}
	for i := range heightMap {
		vis[i] = make([]bool, n)
		for j, height := range heightMap[i] {
			if i == 0 || i == m-1 || j == 0 || j == n-1 {
				heap.Push(pq, &Cell{i, j, height})
				vis[i][j] = true
			}
		}
	}
	out := 0
	dirs := []int{-1, 0, 1, 0, -1}
	for pq.Len() > 0 {
		cell := heap.Pop(pq).(*Cell)
		height := cell.h
		for k := 0; k < 4; k++ {
			x, y := cell.x+dirs[k], cell.y+dirs[k+1]
			if 0 <= x && x < m && 0 <= y && y < n && !vis[x][y] {
				if heightMap[x][y] < height {
					out += height - heightMap[x][y]
				}
				vis[x][y] = true
				heap.Push(pq, &Cell{x, y, Max(height, heightMap[x][y])})
			}
		}
	}
	return out
}

type Cell struct{ x, y, h int }
type PriorityQueue []*Cell

func (h PriorityQueue) Len() int            { return len(h) }
func (h PriorityQueue) Less(i, j int) bool  { return h[i].h < h[j].h }
func (h PriorityQueue) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *PriorityQueue) Push(x interface{}) { *h = append(*h, x.(*Cell)) }
func (h *PriorityQueue) Pop() interface{} {
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
