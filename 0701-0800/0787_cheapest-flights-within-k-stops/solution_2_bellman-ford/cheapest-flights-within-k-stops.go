package main

import "math"

// Tags:
// Bellman-Ford
// Dynamic Programming

func findCheapestPrice(n int, flights [][]int, src int, dst int, K int) int {
	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, K+1)
		for j := range f[i] {
			f[i][j] = math.MaxInt32
		}
	}

	for k := 0; k <= K; k++ {
		f[src][k] = 0
	}

	for _, flight := range flights {
		u, v, price := flight[0], flight[1], flight[2]
		if u == src {
			f[v][0] = price
		}
	}

	for i := 1; i <= K; i++ {
		for _, flight := range flights {
			u, v, price := flight[0], flight[1], flight[2]
			if f[u][i-1] < math.MaxInt32 && f[u][i-1]+price < f[v][i] {
				f[v][i] = f[u][i-1] + flight[2]
			}
		}
	}

	if f[dst][K] == math.MaxInt32 {
		return -1
	}
	return f[dst][K]
}
