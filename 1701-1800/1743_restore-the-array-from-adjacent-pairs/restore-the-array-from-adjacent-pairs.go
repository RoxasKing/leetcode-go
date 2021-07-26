package main

import "bytes"

// Tags:
// Hash

const base int = 1e5

func restoreArray(adjacentPairs [][]int) []int {
	g := [base<<1 + 1][2]int{}
	idx := [base<<1 + 1]byte{}
	for _, p := range adjacentPairs {
		u, v := p[0]+base, p[1]+base
		g[u][idx[u]] = v
		idx[u]++
		g[v][idx[v]] = u
		idx[v]++
	}

	n := len(adjacentPairs) + 1
	out := make([]int, n)
	out[0] = bytes.IndexByte(idx[:], 1) - base
	out[1] = g[out[0]+base][0] - base

	for i := 2; i < n; i++ {
		arr := g[out[i-1]+base]
		if out[i-2] == arr[0]-base {
			out[i] = arr[1] - base
		} else {
			out[i] = arr[0] - base
		}
	}

	return out
}
