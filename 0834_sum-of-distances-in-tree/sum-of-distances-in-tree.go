package main

import "fmt"

/*
  给定一个无向、连通的树。树中有 N 个标记为 0...N-1 的节点以及 N-1 条边 。
  第 i 条边连接节点 edges[i][0] 和 edges[i][1] 。
  返回一个表示节点 i 与其他所有节点距离之和的列表 ans。

  示例 1:
    输入: N = 6, edges = [[0,1],[0,2],[2,3],[2,4],[2,5]]
    输出: [8,12,6,10,10,10]
    解释:
    如下为给定的树的示意图：
        0
       / \
      1   2
         /|\
        3 4 5

    我们可以计算出 dist(0,1) + dist(0,2) + dist(0,3) + dist(0,4) + dist(0,5)
	也就是 1 + 1 + 2 + 2 + 2 = 8。 因此，answer[0] = 8，以此类推。

  说明: 1 <= N <= 10000

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/sum-of-distances-in-tree
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Dynamic Programming
func sumOfDistancesInTree(N int, edges [][]int) []int {
	graph := make([][]int, N)
	for _, e := range edges {
		u, v := e[0], e[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	sz := make([]int, N)
	dp := make([]int, N)
	var dfs func(int, int)
	dfs = func(u, f int) {
		sz[u] = 1
		for _, v := range graph[u] {
			if v == f {
				continue
			}
			dfs(v, u)
			dp[u] += dp[v] + sz[v]
			sz[u] += sz[v]
		}
	}
	dfs(0, -1)

	out := make([]int, N)
	dfs = func(u, f int) {
		out[u] = dp[u]
		fmt.Println(u, graph[u])
		for _, v := range graph[u] {
			if v == f {
				continue
			}
			pu, pv := dp[u], dp[v]
			su, sv := sz[u], sz[v]

			dp[u] -= dp[v] + sz[v]
			sz[u] -= sz[v]
			dp[v] += dp[u] + sz[u]
			sz[v] += sz[u]

			dfs(v, u)

			dp[u], dp[v] = pu, pv
			sz[u], sz[v] = su, sv
		}
	}
	dfs(0, -1)
	return out
}
