package main

import "fmt"

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
