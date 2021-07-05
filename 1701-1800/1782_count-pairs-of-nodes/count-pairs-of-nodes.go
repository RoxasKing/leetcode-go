package main

import "sort"

// Graph + Two Pointers
func countPairs(n int, edges [][]int, queries []int) []int {
	deg := make([]int, n+1)  // count node appear times
	cnt := make(map[int]int) // count edge appear times
	for _, e := range edges {
		deg[e[0]]++
		deg[e[1]]++
		if e[0] > e[1] {
			e[0], e[1] = e[1], e[0]
		}
		cnt[e[0]+(n+1)*e[1]]++
	}

	sortDeg := make([]int, n+1)
	copy(sortDeg, deg)
	sort.Ints(sortDeg)

	out := make([]int, len(queries))
	for i, q := range queries {
		for j, k := 1, n; j <= n; j++ {
			for k >= 1 && sortDeg[j]+sortDeg[k] > q {
				k--
			}
			out[i] += n - Max(j, k)
		}
		for e, t := range cnt {
			j, k := e%(n+1), e/(n+1)
			if deg[j]+deg[k] > q && deg[j]+deg[k]-t <= q {
				out[i]--
			}
		}
	}
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
