package main

// Tags:
// Union-Find

func gardenNoAdj(n int, paths [][]int) []int {
	g := make([][]int, n)
	for _, p := range paths {
		u, v := p[0]-1, p[1]-1
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}
	out := make([]int, n)
	for i := 0; i < n; i++ {
		inValid := [5]bool{}
		for _, v := range g[i] {
			inValid[out[v]] = true
		}
		for k := 1; k <= 4; k++ {
			if !inValid[k] {
				out[i] = k
				break
			}
		}
	}
	return out
}
