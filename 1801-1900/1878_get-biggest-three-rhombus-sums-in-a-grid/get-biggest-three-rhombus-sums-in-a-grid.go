package main

import (
	"container/heap"
)

// Tags:
// Priority Queue

func getBiggestThree(grid [][]int) []int {
	h := MaxHeap{}
	m, n := len(grid), len(grid[0])
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			heap.Push(&h, grid[r][c])
			for i := 1; c-i >= 0 && c+i < n && r+2*i < m; i++ {
				sum := grid[r][c] + grid[r+i][c-i] + grid[r+i][c+i] + grid[r+2*i][c]
				for j := 1; j < i; j++ {
					sum += grid[r+j][c-j] + grid[r+j][c+j] + grid[r+2*i-j][c-j] + grid[r+2*i-j][c+j]
				}
				heap.Push(&h, sum)
			}
		}
	}
	set := map[int]bool{}
	out := make([]int, 0, 3)
	for len(out) < 3 && h.Len() > 0 {
		num := heap.Pop(&h).(int)
		if !set[num] {
			out = append(out, num)
			set[num] = true
		}
	}
	return out
}

type MaxHeap []int

func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() interface{} {
	top := h.Len() - 1
	out := (*h)[top]
	*h = (*h)[:top]
	return out
}
