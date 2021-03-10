package main

import (
	"container/heap"
)

/*
  There is an undirected weighted connected graph. You are given a positive integer n which denotes that the graph has n nodes labeled from 1 to n, and an array edges where each edges[i] = [ui, vi, weighti] denotes that there is an edge between nodes ui and vi with weight equal to weighti.

  A path from node start to node end is a sequence of nodes [z0, z1, z2, ..., zk] such that z0 = start and zk = end and there is an edge between zi and zi+1 where 0 <= i <= k-1.

  The distance of a path is the sum of the weights on the edges of the path. Let distanceToLastNode(x) denote the shortest distance of a path between node n and node x. A restricted path is a path that also satisfies that distanceToLastNode(zi) > distanceToLastNode(zi+1) where 0 <= i <= k-1.

  Return the number of restricted paths from node 1 to node n. Since that number may be too large, return it modulo 10^9 + 7.

  Example 1:
    Input: n = 5, edges = [[1,2,3],[1,3,3],[2,3,1],[1,4,2],[5,2,2],[3,5,1],[5,4,10]]
    Output: 3
    Explanation: Each circle contains the node number in black and its distanceToLastNode value in blue. The three restricted paths are:
    1) 1 --> 2 --> 5
    2) 1 --> 2 --> 3 --> 5
    3) 1 --> 3 --> 5

  Example 2:
    Input: n = 7, edges = [[1,3,1],[4,1,2],[7,3,4],[2,5,3],[5,6,1],[6,7,2],[7,5,3],[2,6,4]]
    Output: 1
    Explanation: Each circle contains the node number in black and its distanceToLastNode value in blue. The only restricted path is 1 --> 3 --> 7.

  Constraints:
    1. 1 <= n <= 2 * 10^4
    2. n - 1 <= edges.length <= 4 * 10^4
    3. edges[i].length == 3
    4. 1 <= ui, vi <= n
    5. ui != vi
    6. 1 <= weighti <= 10^5
    7. There is at most one edge between any two nodes.
    8. There is at least one path between any two nodes.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/number-of-restricted-paths-from-first-to-last-node
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

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
