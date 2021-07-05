package main

import (
	"container/heap"
)

// Tags:
// Graph + Priority Queue + Dynamic Programming
func countRestrictedPaths(n int, edges [][]int) int {
	dict := make([][][2]int, n+1)
	for _, e := range edges {
		dict[e[0]] = append(dict[e[0]], [2]int{e[1], e[2]})
		dict[e[1]] = append(dict[e[1]], [2]int{e[0], e[2]})
	}

	// 节点i到节点n的最小权重距离
	dist := make([]int, n+1)
	for i := range dist {
		dist[i] = 1<<31 - 1
	}
	dist[n] = 0

	// 保存[u, v, weight]的队列
	q := [][3]int{}
	// 先将和节点n直接连接的边加入队列中
	for _, e := range dict[n] {
		q = append(q, [3]int{e[0], n, e[1]})
	}
	// 计算节点1~n-1到节点n的最小权重距离
	for len(q) > 0 {
		e := q[0]
		q = q[1:]
		u, v, w := e[0], e[1], e[2]
		if dist[u] <= dist[v]+w {
			continue
		}
		// 更新节点u到节点n的最小权重距离
		dist[u] = dist[v] + w
		for _, e := range dict[u] {
			if e[0] == v {
				continue
			}
			q = append(q, [3]int{e[0], u, e[1]})
		}
	}

	// 记录从节点1到节点n的限制路径数
	dp := make([]int, n+1)
	dp[1] = 1

	// 记录已经走过过的路径
	walked := make(map[[2]int]bool)

	// 将边[u, v]按v的最小权重距离从大到小排列的优先队列
	pq := &PriorityQueue{dists: dist}
	// 先将和节点1直接连接的路径加入到队列中
	for _, e := range dict[1] {
		v := e[0]
		if dist[1] > dist[v] {
			heap.Push(pq, [2]int{1, v})
		}
	}

	for pq.Len() > 0 {
		e := heap.Pop(pq).([2]int)
		u, v := e[0], e[1]

		if walked[[2]int{u, v}] {
			continue
		}
		walked[[2]int{u, v}] = true

		// 累加限制路径的数量
		dp[v] = (dp[v] + dp[u]) % (1e9 + 7)

		for _, e := range dict[v] {
			edge := [2]int{v, e[0]}
			if dist[v] <= dist[e[0]] || walked[edge] {
				continue
			}
			heap.Push(pq, edge)
		}
	}

	return dp[n]
}

type PriorityQueue struct {
	edges [][2]int // [u, v] pair list
	dists []int
}

func (pq PriorityQueue) Len() int { return len(pq.edges) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq.dists[pq.edges[i][1]] > pq.dists[pq.edges[j][1]]
}
func (pq PriorityQueue) Swap(i, j int)       { pq.edges[i], pq.edges[j] = pq.edges[j], pq.edges[i] }
func (pq *PriorityQueue) Push(x interface{}) { pq.edges = append(pq.edges, x.([2]int)) }
func (pq *PriorityQueue) Pop() interface{} {
	last := pq.Len() - 1
	out := pq.edges[last]
	pq.edges = pq.edges[:last]
	return out
}
