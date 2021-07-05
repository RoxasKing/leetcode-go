package main

import "math"

// Dynamic Programming

func minSkips(dist []int, speed int, hoursBefore int) int {
	n := len(dist)
	f := make([][]float64, n)
	for i := range f {
		f[i] = make([]float64, n)
		for j := range f[i] {
			f[i][j] = 1e9
		}
	}
	f[0][0] = float64(dist[0]) / float64(speed)

	for i := 1; i < n; i++ {
		cost := float64(dist[i]) / float64(speed)
		for t := 0; t < n; t++ {
			if f[i-1][t] == 1e9 {
				continue
			}
			f[i][t] = math.Min(f[i][t], math.Ceil(f[i-1][t]-1e-9)+cost)
			f[i][t+1] = math.Min(f[i][t+1], f[i-1][t]+cost)
		}
	}

	for i := range f[n-1] {
		if f[n-1][i] <= float64(hoursBefore) {
			return i
		}
	}
	return -1
}
