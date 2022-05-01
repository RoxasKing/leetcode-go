package main

// Difficulty:
// Medium

// Tags:
// Backtracking
// Hash

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	type next struct {
		k string
		v float64
	}
	h := map[string][]*next{}
	for i, e := range equations {
		h[e[0]] = append(h[e[0]], &next{e[1], values[i]})
		h[e[1]] = append(h[e[1]], &next{e[0], 1 / values[i]})
	}
	var vis map[string]struct{}
	var backtrack func(string, string, float64) float64
	backtrack = func(a, b string, v float64) float64 {
		vis[a] = struct{}{}
		for _, x := range h[a] {
			if x.k == b {
				return v * x.v
			}
			if _, ok := vis[x.k]; ok {
				continue
			}
			if res := backtrack(x.k, b, v*x.v); res != -1.0 {
				return res
			}
		}
		delete(vis, a)
		return -1.0
	}
	out := make([]float64, len(queries))
	for i, q := range queries {
		vis = map[string]struct{}{}
		out[i] = backtrack(q[0], q[1], 1)
	}
	return out
}
