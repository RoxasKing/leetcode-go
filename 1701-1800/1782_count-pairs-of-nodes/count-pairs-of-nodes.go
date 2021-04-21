package main

import "sort"

/*
  You are given an undirected graph represented by an integer n, which is the number of nodes, and edges, where edges[i] = [ui, vi] which indicates that there is an undirected edge between ui and vi. You are also given an integer array queries.

  The answer to the jth query is the number of pairs of nodes (a, b) that satisfy the following conditions:
    1. a < b
    2. cnt is strictly greater than queries[j], where cnt is the number of edges incident to a or b.
  Return an array answers such that answers.length == queries.length and answers[j] is the answer of the jth query.

  Note that there can be repeated edges.

  Example 1:
    Input: n = 4, edges = [[1,2],[2,4],[1,3],[2,3],[2,1]], queries = [2,3]
    Output: [6,5]
    Explanation: The number of edges incident to at least one of each pair is shown above.

  Example 2:
    Input: n = 5, edges = [[1,5],[1,5],[3,4],[2,5],[1,3],[5,1],[2,3],[2,5]], queries = [1,2,3,4,5]
    Output: [10,10,9,8,6]

  Constraints:
    1. 2 <= n <= 2 * 10^4
    2. 1 <= edges.length <= 10^5
    3. 1 <= ui, vi <= n
    4. ui != vi
    5. 1 <= queries.length <= 20
    6. 0 <= queries[j] < edges.length

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/count-pairs-of-nodes
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Graph + Two Pointers
func countPairs(n int, edges [][]int, queries []int) []int {
	deg := make([]int, n+1)  // count node appear times
	cnt := make(map[int]int) // count edge appear times
	for _, e := range edges {
		deg[e[0]]++
		deg[e[1]]++
		if e[0] > e[1] {
			e[0], e[1] = e[1], e[0]
		}
		cnt[e[0]+(n+1)*e[1]]++
	}

	sortDeg := make([]int, n+1)
	copy(sortDeg, deg)
	sort.Ints(sortDeg)

	out := make([]int, len(queries))
	for i, q := range queries {
		for j, k := 1, n; j <= n; j++ {
			for k >= 1 && sortDeg[j]+sortDeg[k] > q {
				k--
			}
			out[i] += n - Max(j, k)
		}
		for e, t := range cnt {
			j, k := e%(n+1), e/(n+1)
			if deg[j]+deg[k] > q && deg[j]+deg[k]-t <= q {
				out[i]--
			}
		}
	}
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
