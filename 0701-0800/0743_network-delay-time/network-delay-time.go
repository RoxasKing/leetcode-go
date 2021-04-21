package main

import "container/heap"

/*
  You are given a network of n nodes, labeled from 1 to n. You are also given times, a list of travel times as directed edges times[i] = (ui, vi, wi), where ui is the source node, vi is the target node, and wi is the time it takes for a signal to travel from source to target.

  We will send a signal from a given node k. Return the time it takes for all the n nodes to receive the signal. If it is impossible for all the n nodes to receive the signal, return -1.

  Example 1:
    Input: times = [[2,1,1],[2,3,1],[3,4,1]], n = 4, k = 2
    Output: 2

  Example 2:
    Input: times = [[1,2,1]], n = 2, k = 1
    Output: 1

  Example 3:
    Input: times = [[1,2,1]], n = 2, k = 2
    Output: -1

  Constraints:
    1. 1 <= k <= n <= 100
    2. 1 <= times.length <= 6000
    3. times[i].length == 3
    4. 1 <= ui, vi <= n
    5. ui != vi
    6. 0 <= wi <= 100
    7. All the pairs (ui, vi) are unique. (i.e., no multiple edges.)

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/network-delay-time
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

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
