package main

import (
	"sort"
)

// Tags:
// Graph

func minTrioDegree(n int, edges [][]int) int {
	linked := make([][]bool, n+1)
	for i := range linked {
		linked[i] = make([]bool, n+1)
	}
	indeg := make([]int, n+1)
	es := make([][]int, n+1)
	for _, e := range edges {
		if e[0] > e[1] {
			e[0], e[1] = e[1], e[0]
		}
		u, v := e[0], e[1]
		linked[u][v] = true
		indeg[u]++
		indeg[v]++
		es[u] = append(es[u], v)
	}

	out := 1<<31 - 1

	for a := 1; a <= n; a++ {
		sort.Ints(es[a])
		for j, b := range es[a] {
			for _, c := range es[a][j+1:] {
				if !linked[b][c] {
					continue
				}
				out = Min(out, indeg[a]+indeg[b]+indeg[c]-6)
			}
		}
	}

	if out == 1<<31-1 {
		return -1
	}

	return out
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
