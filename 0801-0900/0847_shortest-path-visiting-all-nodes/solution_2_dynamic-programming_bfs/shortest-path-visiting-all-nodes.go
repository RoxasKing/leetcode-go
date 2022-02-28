package main

import "runtime/debug"

// Difficulty:
// Hard

// Tags:
// Dynamic-programming
// BFS
// Bit Manipulation
// Bitmask

func init() { debug.SetGCPercent(-1) }

func shortestPathLength(graph [][]int) int {
	type state struct{ node, step, mask int }
	q := []state{}
	n := len(graph)
	vis := make([][]bool, n)
	for i := 0; i < n; i++ {
		q = append(q, state{i, 0, 1 << i})
		vis[i] = make([]bool, 1<<n)
		vis[i][1<<i] = true
	}
	for ; q[0].mask != 1<<n-1; q = q[1:] {
		u := q[0]
		for _, v := range graph[u.node] {
			if mask := u.mask | (1 << v); !vis[v][mask] {
				vis[v][mask] = true
				q = append(q, state{v, u.step + 1, mask})
			}
		}
	}
	return q[0].step
}
