package main

// Dfficulty:
// Medium

// Tags:
// BFS

func findMinHeightTrees(n int, edges [][]int) []int {
	d := make([]int, n)
	g := make([][]int, n)
	for _, e := range edges {
		a, b := e[0], e[1]
		d[a]++
		d[b]++
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}
	var q []int
	for i := 0; i < n; i++ {
		if d[i] <= 1 {
			q = append(q, i)
		}
	}
	var out []int
	for len(q) > 0 {
		var t []int
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
