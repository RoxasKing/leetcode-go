package main

// Difficulty:
// Medium

// Tags:
// DFS

func countHighestScoreNodes(parents []int) int {
	n := len(parents)
	g := make([][]int, n)
	for i, p := range parents {
		if p == -1 {
			continue
		}
		g[p] = append(g[p], i)
	}
	c := make([]int, n)
	var dfs func(int)
	dfs = func(u int) {
		c[u] = 1
		for _, v := range g[u] {
			dfs(v)
			c[u] += c[v]
		}
	}
	dfs(0)
	out, cnt := 1, 1
	for _, v := range g[0] {
		out *= c[v]
	}
	for i := 1; i < n; i++ {
		score := c[0] - c[i]
		for _, v := range g[i] {
			score *= c[v]
		}
		if out == score {
			cnt++
		} else if out < score {
			out, cnt = score, 1
		}
	}
	return cnt
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
