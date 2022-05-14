package main

// Difficulty:
// Medium

// Tags:
// Graph
// BFS

func networkDelayTime(times [][]int, n int, k int) int {
	g := make([][][]int, n+1)
	for _, e := range times {
		u, v, w := e[0], e[1], e[2]
		g[u] = append(g[u], []int{v, w})
	}
	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		f[i] = 1<<31 - 1
	}
	f[k] = 0
	for q := []int{k}; len(q) > 0; q = q[1:] {
		u := q[0]
		for _, e := range g[u] {
			v, w := e[0], e[1]
			if f[v] <= f[u]+w {
				continue
			}
			f[v] = f[u] + w
			q = append(q, v)
		}
	}
	o := 0
	for i := 1; i <= n; i++ {
		if f[i] == 1<<31-1 {
			return -1
		}
		if o < f[i] {
			o = f[i]
		}
	}
	return o
}
