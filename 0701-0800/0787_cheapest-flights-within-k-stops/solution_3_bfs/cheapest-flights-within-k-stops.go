package main

import "math"

// Tags:
// BFS
// Dynamic Programming

func findCheapestPrice(n int, flights [][]int, src int, dst int, K int) int {
	g := make([][]int, n)
	for i := range g {
		g[i] = make([]int, n)
	}
	for _, flight := range flights {
		u, v, price := flight[0], flight[1], flight[2]
		g[u][v] = price
	}

	f := make([]int, n)
	for i := range f {
		f[i] = math.MaxInt32
	}
	f[src] = 0

	q := [][]int{{src, 0}}
	for ; K >= 0; K-- {
		m := len(q)
		for _, p := range q {
			u, price := p[0], p[1]
			for v := 0; v < n; v++ {
				if g[u][v] == 0 || price+g[u][v] > f[v] {
					continue
				}
				f[v] = price + g[u][v]
				q = append(q, []int{v, f[v]})
			}
		}
		q = q[m:]
	}

	if f[dst] == math.MaxInt32 {
		return -1
	}
	return f[dst]
}
