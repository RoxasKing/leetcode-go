package main

/*
  A tree is an undirected graph in which any two vertices are connected by exactly one path. In other words, any connected graph without simple cycles is a tree.

  Given a tree of n nodes labelled from 0 to n - 1, and an array of n - 1 edges where edges[i] = [ai, bi] indicates that there is an undirected edge between the two nodes ai and bi in the tree, you can choose any node of the tree as the root. When you select a node x as the root, the result tree has height h. Among all possible rooted trees, those with minimum height (i.e. min(h))  are called minimum height trees (MHTs).

  Return a list of all MHTs' root labels. You can return the answer in any order.

  The height of a rooted tree is the number of edges on the longest downward path between the root and a leaf.

  Example 1:
    Input: n = 4, edges = [[1,0],[1,2],[1,3]]
    Output: [1]
    Explanation: As shown, the height of the tree is 1 when the root is the node with label 1 which is the only MHT.

  Example 2:
    Input: n = 6, edges = [[3,0],[3,1],[3,2],[3,4],[5,4]]
    Output: [3,4]

  Example 3:
    Input: n = 1, edges = []
    Output: [0]

  Example 4:
    Input: n = 2, edges = [[0,1]]
    Output: [0,1]

  Constraints:
    1. 1 <= n <= 2 * 10^4
    2. edges.length == n - 1
    3. 0 <= ai, bi < n
    4. ai != bi
    5. All the pairs (ai, bi) are distinct.
    6. The given input is guaranteed to be a tree and there will be no repeated edges.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/minimum-height-trees
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// BFS

func findMinHeightTrees(n int, edges [][]int) []int {
	indeg := make([]int, n)
	next := make([][]int, n)
	for _, e := range edges {
		a, b := e[0], e[1]
		indeg[a]++
		indeg[b]++
		next[a] = append(next[a], b)
		next[b] = append(next[b], a)
	}

	q := []int{}
	for i := 0; i < n; i++ {
		if indeg[i] <= 1 {
			q = append(q, i)
		}
	}

	var out []int
	for len(q) > 0 {
		out = q
		nq := []int{}
		for _, u := range q {
			for _, v := range next[u] {
				indeg[v]--
				if indeg[v] == 1 {
					nq = append(nq, v)
				}
			}
		}
		q = nq
	}
	return out
}
