package main

import (
	"container/heap"
)

/*
  There are n cities connected by m flights. Each flight starts from city u and arrives at v with a price w.

  Now given all the cities and flights, together with starting city src and the destination dst, your task is to find the cheapest price from src to dst with up to k stops. If there is no such route, output -1.

  Example 1:
    Input:
    n = 3, edges = [[0,1,100],[1,2,100],[0,2,500]]
    src = 0, dst = 2, k = 1
    Output: 200
    Explanation:
    The graph looks like this:
    The cheapest price from city 0 to city 2 with at most 1 stop costs 200, as marked red in the picture.

  Example 2:
    Input:
    n = 3, edges = [[0,1,100],[1,2,100],[0,2,500]]
    src = 0, dst = 2, k = 0
    Output: 500
    Explanation:
    The graph looks like this:

  The cheapest price from city 0 to city 2 with at most 0 stop costs 500, as marked blue in the picture.

  Constraints:
    1. The number of nodes n will be in range [1, 100], with nodes labeled from 0 to n - 1.
    2. The size of flights will be in range [0, n * (n - 1) / 2].
    3. The format of each flight will be (src, dst, price).
    4. The price of each flight will be in the range [1, 10000].
    5. k is in the range of [0, n - 1].
    6. There will not be any duplicated flights or self cycles.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/cheapest-flights-within-k-stops
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

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
