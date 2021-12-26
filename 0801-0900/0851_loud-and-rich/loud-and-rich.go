package main

// Difficulty:
// Medium

// Tags:
// Topological Sort

func loudAndRich(richer [][]int, quiet []int) []int {
	n := len(quiet)
	g := make([][]int, n)
	d := make([]int, n)
	for i := range richer {
		a, b := richer[i][0], richer[i][1]
		g[a] = append(g[a], b)
		d[b]++
	}
	q := []int{}
	for i := 0; i < n; i++ {
		if d[i] == 0 {
			q = append(q, i)
		}
	}
	out := make([]int, n)
	for i := range out {
		out[i] = i
	}
	for ; len(q) > 0; q = q[1:] {
		a := q[0]
		for _, b := range g[a] {
			if quiet[out[a]] < quiet[out[b]] {
				out[b] = out[a]
			}
			d[b]--
			if d[b] == 0 {
				q = append(q, b)
			}
		}
	}
	return out
}
