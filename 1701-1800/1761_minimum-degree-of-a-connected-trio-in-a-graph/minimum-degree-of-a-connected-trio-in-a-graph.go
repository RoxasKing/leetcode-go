package main

import (
	"sort"
)

/*
  You are given an undirected graph. You are given an integer n which is the number of nodes in the graph and an array edges, where each edges[i] = [ui, vi] indicates that there is an undirected edge between ui and vi.

  A connected trio is a set of three nodes where there is an edge between every pair of them.

  The degree of a connected trio is the number of edges where one endpoint is in the trio, and the other is not.

  Return the minimum degree of a connected trio in the graph, or -1 if the graph has no connected trios.

  Example 1:
    Input: n = 6, edges = [[1,2],[1,3],[3,2],[4,1],[5,2],[3,6]]
    Output: 3
    Explanation: There is exactly one trio, which is [1,2,3]. The edges that form its degree are bolded in the figure above.

  Example 2:
    Input: n = 7, edges = [[1,3],[4,1],[4,3],[2,5],[5,6],[6,7],[7,5],[2,6]]
    Output: 0
    Explanation: There are exactly three trios:
      1) [1,4,3] with degree 0.
      2) [2,5,6] with degree 2.
      3) [5,6,7] with degree 2.

  Constraints:
    1. 2 <= n <= 400
    2. edges[i].length == 2
    3. 1 <= edges.length <= n * (n-1) / 2
    4. 1 <= ui, vi <= n
    5. ui != vi
    6. There are no repeated edges.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/minimum-degree-of-a-connected-trio-in-a-graph
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// Graph

func minTrioDegree(n int, edges [][]int) int {
	linked := make([][]bool, n+1)
	for i := range linked {
		linked[i] = make([]bool, n+1)
	}
	indeg := make([]int, n+1)
	es := make([][]int, n+1)
	for _, e := range edges {
		if e[0] > e[1] {
			e[0], e[1] = e[1], e[0]
		}
		u, v := e[0], e[1]
		linked[u][v] = true
		indeg[u]++
		indeg[v]++
		es[u] = append(es[u], v)
	}

	out := 1<<31 - 1

	for a := 1; a <= n; a++ {
		sort.Ints(es[a])
		for j, b := range es[a] {
			for _, c := range es[a][j+1:] {
				if !linked[b][c] {
					continue
				}
				out = Min(out, indeg[a]+indeg[b]+indeg[c]-6)
			}
		}
	}

	if out == 1<<31-1 {
		return -1
	}

	return out
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
