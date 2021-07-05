package main

import (
	"container/heap"
)

// Tags:
// Dijkstra + Priority Queue
func findCheapestPrice(n int, flights [][]int, src int, dst int, K int) int {
	graph := make([][]int, n)
	for i := range graph {
		graph[i] = make([]int, n)
	}
	for _, f := range flights {
		graph[f[0]][f[1]] = f[2]
	}

	pq := &PriorityQueue{}
	heap.Push(pq, [3]int{src, 0, K + 1})
	for pq.Len() > 0 {
		p := heap.Pop(pq).([3]int)
		v, price, k := p[0], p[1], p[2]

		if v == dst {
			return price
		}

		if k == 0 {
			continue
		}

		for i := 0; i < n; i++ {
			if graph[v][i] == 0 {
				continue
			}
			heap.Push(pq, [3]int{i, price + graph[v][i], k - 1})
		}
	}

	return -1
}

type PriorityQueue [][3]int // [dst, price, k]

func (pq PriorityQueue) Len() int            { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool  { return pq[i][1] < pq[j][1] }
func (pq PriorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) { *pq = append(*pq, x.([3]int)) }
func (pq *PriorityQueue) Pop() interface{} {
	top := len(*pq) - 1
	out := (*pq)[top]
	*pq = (*pq)[:top]
	return out
}

// Bellman-Ford + Dynamic Programming
func findCheapestPrice2(n int, flights [][]int, src int, dst int, K int) int {
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, K+1)
		for j := range dp[i] {
			dp[i][j] = 1<<31 - 1
		}
	}

	// 自己离自己最短距离永远是0
	for i := 0; i <= K; i++ {
		dp[src][i] = 0
	}

	// 与src直接连接的节点
	for _, f := range flights {
		if f[0] == src {
			dp[f[1]][0] = f[2]
		}
	}

	// K 次 松弛操作(k个中转点)
	for i := 1; i <= K; i++ {
		for _, f := range flights {
			u, v := f[0], f[1]
			if dp[u][i-1] < 1<<31-1 && dp[u][i-1]+f[2] < dp[v][i] {
				dp[v][i] = dp[u][i-1] + f[2]
			}
		}
	}

	if dp[dst][K] == 1<<31-1 {
		return -1
	}
	return dp[dst][K]
}
