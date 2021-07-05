package main

import "container/heap"

// Dijkstra's Algorithm + Priority Queue
func networkDelayTime(times [][]int, N int, K int) int {
	edges := make([][][2]int, N+1)
	for _, t := range times {
		u, v, w := t[0], t[1], t[2]
		edges[u] = append(edges[u], [2]int{v, w})
	}

	mark := make([]int, N+1)
	for i := 1; i <= N; i++ {
		mark[i] = 1<<31 - 1
	}
	mark[K] = 0

	pq := PriorityQueue{}
	heap.Push(&pq, [2]int{K, 0})

	for pq.Len() > 0 {
		u := heap.Pop(&pq).([2]int)[0]
		for _, e := range edges[u] {
			v, w := e[0], e[1]
			if mark[u]+w < mark[v] {
				mark[v] = mark[u] + w
				heap.Push(&pq, e)
			}
		}
	}

	out := 0
	for i := 1; i <= N; i++ {
		if mark[i] == 1<<31-1 {
			return -1
		}
		out = Max(out, mark[i])
	}
	return out
}

type PriorityQueue [][2]int

func (pq PriorityQueue) Len() int            { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool  { return pq[i][1] < pq[j][1] }
func (pq PriorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) { *pq = append(*pq, x.([2]int)) }
func (pq *PriorityQueue) Pop() interface{} {
	top := pq.Len() - 1
	out := (*pq)[top]
	*pq = (*pq)[:top]
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// DFS
func networkDelayTime2(times [][]int, N int, K int) int {
	graph := make(map[int][][]int)
	for _, t := range times {
		graph[t[0]] = append(graph[t[0]], []int{t[1], t[2]})
	}
	dist := make(map[int]int)
	for i := 1; i <= N; i++ {
		dist[i] = 1<<31 - 1
	}
	dfsNetworkDelayTime(&graph, &dist, K, 0)
	var out int
	for _, v := range dist {
		if v == 1<<31-1 {
			return -1
		}
		out = Max(out, v)
	}
	return out
}

func dfsNetworkDelayTime(graph *map[int][][]int, dist *map[int]int, node int, cur int) {
	if cur >= (*dist)[node] {
		return
	}
	(*dist)[node] = cur
	if _, ok := (*graph)[node]; ok {
		for _, elem := range (*graph)[node] {
			dfsNetworkDelayTime(graph, dist, elem[0], cur+elem[1])
		}
	}
}
