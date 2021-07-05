package main

// Tags:
// DFS
func numWays(n int, relation [][]int, k int) int {
	edges := make([][]int, n)
	for _, r := range relation {
		edges[r[0]] = append(edges[r[0]], r[1])
	}
	out := 0
	dfs(edges, n, k, 0, 0, &out)
	return out
}

func dfs(edges [][]int, n, k, i, cnt int, out *int) {
	if cnt == k {
		if i == n-1 {
			*out++
		}
		return
	}
	for _, next := range edges[i] {
		dfs(edges, n, k, next, cnt+1, out)
	}
}

// BFS
func numWays2(n int, relation [][]int, k int) int {
	edges := make(map[int][]int)
	for _, r := range relation {
		edges[r[0]] = append(edges[r[0]], r[1])
	}
	out := 0
	q := [][]int{{0, 0}}
	for len(q) > 0 {
		e := q[0]
		q = q[1:]
		num, cnt := e[0], e[1]
		if cnt == k {
			if num == n-1 {
				out++
			}
			continue
		}
		for _, next := range edges[num] {
			q = append(q, []int{next, cnt + 1})
		}
	}
	return out
}

// Dynamic Programming
func numWays3(n int, relation [][]int, k int) int {
	dp0 := make([]int, n)
	dp0[0] = 1
	for i := 0; i < k; i++ {
		dp := make([]int, n)
		for _, r := range relation {
			dp[r[1]] += dp0[r[0]]
		}
		dp0 = dp
	}
	return dp0[n-1]
}
