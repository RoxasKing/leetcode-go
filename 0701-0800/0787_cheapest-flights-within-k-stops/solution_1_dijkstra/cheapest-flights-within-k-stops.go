package main

import "container/heap"

// Tags:
// Dijkstra
// Priority Queue
// Dynamic Programming

func findCheapestPrice(n int, flights [][]int, src int, dst int, K int) int {
	g := make([][]int, n)
	for i := range g {
		g[i] = make([]int, n)
	}
	for _, f := range flights {
		g[f[0]][f[1]] = f[2]
	}

	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, K+1)
		for j := range f[i] {
			f[i][j] = 1e9
		}
	}
	for i := 0; i <= K; i++ {
		f[src][i] = 0
	}

	pq := &PriorityQueue{pair{src, 0, K + 1}}
	for pq.Len() > 0 {
		p := heap.Pop(pq).(pair)
		u, price, k := p.u, p.price, p.k

		if u == dst {
			return price
		}

		if k == 0 {
			continue
		}

		for v := 0; v < n; v++ {
			if g[u][v] == 0 || price+g[u][v] >= f[v][k-1] {
				continue
			}
			f[v][k-1] = price + g[u][v]
			heap.Push(pq, pair{v, price + g[u][v], k - 1})
		}
	}

	return -1
}

type pair struct {
	u, price, k int
}

type PriorityQueue []pair

func (pq PriorityQueue) Len() int            { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool  { return pq[i].price < pq[j].price }
func (pq PriorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) { *pq = append(*pq, x.(pair)) }
func (pq *PriorityQueue) Pop() interface{} {
	top := len(*pq) - 1
	out := (*pq)[top]
	*pq = (*pq)[:top]
	return out
}
