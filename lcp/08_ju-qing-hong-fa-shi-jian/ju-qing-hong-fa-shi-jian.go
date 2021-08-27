package main

import (
	"container/heap"
	"sort"
)

// Tags:
// Binary Search
func getTriggerTime(increase [][]int, requirements [][]int) []int {
	m := len(increase)
	for i := 1; i < m; i++ {
		increase[i][0] += increase[i-1][0]
		increase[i][1] += increase[i-1][1]
		increase[i][2] += increase[i-1][2]
	}

	out := make([]int, len(requirements))
	for i, req := range requirements {
		if req[0] == 0 && req[1] == 0 && req[2] == 0 {
			continue
		}
		if req[0] > increase[m-1][0] || req[1] > increase[m-1][1] || req[2] > increase[m-1][2] {
			out[i] = -1
			continue
		}
		out[i] = sort.Search(m, func(i int) bool {
			return increase[i][0] >= req[0] && increase[i][1] >= req[1] && increase[i][2] >= req[2]
		}) + 1
	}

	return out
}

// Priority Queue
func getTriggerTime2(increase [][]int, requirements [][]int) []int {
	n := len(requirements)
	pqC := NewPriorityQueue(requirements, 0)
	pqR := NewPriorityQueue(requirements, 1)
	pqH := NewPriorityQueue(requirements, 2)
	c, r, h := 0, 0, 0
	out := make([]int, n)
	for i, req := range requirements {
		if req[0] > c || req[1] > r || req[2] > h {
			heap.Push(pqC, i)
			heap.Push(pqR, i)
			heap.Push(pqH, i)
			out[i] = -1
		}
	}

	for i, inc := range increase {
		c += inc[0]
		r += inc[1]
		h += inc[2]
		for pqC.Len() > 0 && requirements[pqC.idxs[0]][0] <= c {
			idx := heap.Pop(pqC).(int)
			if out[idx] != -1 {
				continue
			}
			if requirements[idx][1] <= r && requirements[idx][2] <= h {
				out[idx] = i + 1
			}
		}
		for pqR.Len() > 0 && requirements[pqR.idxs[0]][1] <= r {
			idx := heap.Pop(pqR).(int)
			if out[idx] != -1 {
				continue
			}
			if requirements[idx][0] <= c && requirements[idx][2] <= h {
				out[idx] = i + 1
			}
		}
		for pqH.Len() > 0 && requirements[pqH.idxs[0]][2] <= h {
			idx := heap.Pop(pqH).(int)
			if out[idx] != -1 {
				continue
			}
			if requirements[idx][0] <= c && requirements[idx][1] <= r {
				out[idx] = i + 1
			}
		}
	}
	return out
}

type PriorityQueue struct {
	reqs [][]int
	idxs []int
	typ  int
}

func NewPriorityQueue(reqs [][]int, typ int) *PriorityQueue {
	return &PriorityQueue{
		reqs: reqs,
		typ:  typ,
	}
}

func (pq PriorityQueue) Len() int { return len(pq.idxs) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq.reqs[pq.idxs[i]][pq.typ] < pq.reqs[pq.idxs[j]][pq.typ]
}
func (pq PriorityQueue) Swap(i, j int) {
	pq.idxs[i], pq.idxs[j] = pq.idxs[j], pq.idxs[i]
}
func (pq *PriorityQueue) Push(x interface{}) { pq.idxs = append(pq.idxs, x.(int)) }
func (pq *PriorityQueue) Pop() interface{} {
	top := pq.Len() - 1
	out := pq.idxs[top]
	pq.idxs = pq.idxs[:top]
	return out
}
