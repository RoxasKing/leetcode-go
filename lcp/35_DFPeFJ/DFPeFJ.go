package main

import "container/heap"

// Dijkstra's algorithm + Priority Queue
func electricCarPlan(paths [][]int, cnt int, start int, end int, charge []int) int {
	n := len(charge)
	edges := make([][][2]int, n)
	for _, p := range paths {
		u, v, w := p[0], p[1], p[2]
		edges[u] = append(edges[u], [2]int{v, w})
		edges[v] = append(edges[v], [2]int{u, w})
	}

	pq := PriorityQueue{}
	heap.Push(&pq, [3]int{0, start, 0})

	minT := make([][]int, n) // 记录到达不同地点不同电量的最少耗时
	for i := range minT {
		minT[i] = make([]int, cnt+1)
		for j := range minT[i] {
			minT[i][j] = 1<<31 - 1
		}
	}
	minT[start][0] = 0

	for pq.Len() > 0 {
		e := heap.Pop(&pq).([3]int)
		t, u, c := e[0], e[1], e[2] // 耗时， 地点， 电量

		if t > minT[u][c] { // 大于已记录的最佳耗时，跳过
			continue
		}

		if u == end {
			return t
		}

		if c < cnt { // 充电
			newT := t + charge[u]
			if newT < minT[u][c+1] {
				minT[u][c+1] = newT
				heap.Push(&pq, [3]int{newT, u, c + 1})
			}
		}

		for _, next := range edges[u] {
			v, w := next[0], next[1]
			if c >= w && t+w < minT[v][c-w] { // 电量足够，可达，并且比记录的耗时低，更新记录，并入队
				minT[v][c-w] = t + w
				heap.Push(&pq, [3]int{t + w, v, c - w})
			}
		}
	}

	return -1
}

type PriorityQueue [][3]int // 耗时， 地点， 电量

func (pq PriorityQueue) Len() int            { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool  { return pq[i][0] < pq[j][0] }
func (pq PriorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) { *pq = append(*pq, x.([3]int)) }
func (pq *PriorityQueue) Pop() interface{} {
	top := pq.Len() - 1
	out := (*pq)[top]
	*pq = (*pq)[:top]
	return out
}
