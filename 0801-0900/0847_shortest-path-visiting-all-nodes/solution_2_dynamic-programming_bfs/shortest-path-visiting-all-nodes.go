package main

// Tags:
// Dynamic-programming
// BFS

func shortestPathLength(graph [][]int) int {
	n := len(graph)
	seen := make([][]bool, n)
	q := []pair{}
	for i := 0; i < n; i++ {
		seen[i] = make([]bool, 1<<n)
		seen[i][1<<i] = true
		q = append(q, pair{i, 0, 1 << i})
	}
	for ; ; q = q[1:] {
		p := q[0]
		if p.mask == 1<<n-1 {
			return p.step
		}
		for _, v := range graph[p.u] {
			mask := p.mask | (1 << v)
			if !seen[v][mask] {
				seen[v][mask] = true
				q = append(q, pair{v, p.step + 1, mask})
			}
		}
	}
}

type pair struct {
	u, step, mask int
}
