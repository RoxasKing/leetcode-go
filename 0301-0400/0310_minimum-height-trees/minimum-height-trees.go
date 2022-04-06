package main

// Dfficulty:
// Medium

// Tags:
// BFS

func findMinHeightTrees(n int, edges [][]int) []int {
	d := make([]int, n)
	g := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		d[u]++
		d[v]++
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}
	q := []int{}
	for i := 0; i < n; i++ {
		if d[i] <= 1 {
			q = append(q, i)
		}
	}
	out := []int{}
	for len(q) > 0 {
		t := []int{}
		for _, u := range q {
			for _, v := range g[u] {
				if d[v]--; d[v] == 1 {
					t = append(t, v)
				}
			}
		}
		out, q = q, t
	}
	return out
}
