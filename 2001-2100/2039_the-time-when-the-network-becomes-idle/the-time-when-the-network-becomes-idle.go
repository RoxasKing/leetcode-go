package main

// Difficulty:
// Medium

// Tags:
// Graph
// BFS
// Dynamic Programming

func networkBecomesIdle(edges [][]int, patience []int) int {
	n := len(patience)
	g := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}
	f := make([]int, n)
	for i := 1; i < n; i++ {
		f[i] = 1e9
	}
	for q := [][2]int{{0, 0}}; len(q) > 0; q = q[1:] {
		e := q[0]
		u, d := e[0], e[1]+1
		for _, v := range g[u] {
			if f[v] <= d {
				continue
			}
			f[v] = d
			q = append(q, [2]int{v, d})
		}
	}
	out := 0
	for i := 1; i < n; i++ {
		t0 := f[i] << 1
		tt := t0 / patience[i]
		if t0%patience[i] == 0 {
			tt--
		}
		out = Max(out, t0+tt*patience[i])
	}
	return out + 1
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
