package main

/*
  There is a directed graph of n colored nodes and m edges. The nodes are numbered from 0 to n - 1.

  You are given a string colors where colors[i] is a lowercase English letter representing the color of the ith node in this graph (0-indexed). You are also given a 2D array edges where edges[j] = [aj, bj] indicates that there is a directed edge from node aj to node bj.

  A valid path in the graph is a sequence of nodes x1 -> x2 -> x3 -> ... -> xk such that there is a directed edge from xi to xi+1 for every 1 <= i < k. The color value of the path is the number of nodes that are colored the most frequently occurring color along that path.

  Return the largest color value of any valid path in the given graph, or -1 if the graph contains a cycle.

  Example 1:
    Input: colors = "abaca", edges = [[0,1],[0,2],[2,3],[3,4]]
    Output: 3
    Explanation: The path 0 -> 2 -> 3 -> 4 contains 3 nodes that are colored "a" (red in the above image).

  Example 2:
    Input: colors = "a", edges = [[0,0]]
    Output: -1
    Explanation: There is a cycle from 0 to 0.

  Constraints:
    1. n == colors.length
    2. m == edges.length
    3. 1 <= n <= 10^5
    4. 0 <= m <= 10^5
    5. colors consists of lowercase English letters.
    6. 0 <= aj, bj < n

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/largest-color-value-in-a-directed-graph
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Topological Sort + Dynamic Programming
func largestPathValue(colors string, edges [][]int) int {
	if len(edges) == 0 {
		return 1
	}

	n := len(colors)
	nexts := make([][]int, n)
	indeg := make([]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		nexts[u] = append(nexts[u], v)
		indeg[v]++
	}

	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, 26)
	}

	q := []int{}
	for i := 0; i < n; i++ {
		if indeg[i] == 0 && len(nexts[i]) > 0 {
			q = append(q, i)
		}
	}

	out := -1
	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		f[u][colors[u]-'a']++
		if len(nexts[u]) == 0 {
			for i := 0; i < 26; i++ {
				out = Max(out, f[u][i])
			}
		}
		for _, v := range nexts[u] {
			for i := 0; i < 26; i++ {
				f[v][i] = Max(f[v][i], f[u][i])
			}
			indeg[v]--
			if indeg[v] == 0 {
				q = append(q, v)
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
