package main

// Tags:
// Dynamic Programming
// Floyd–Warshall

func shortestPathLength(graph [][]int) int {
	n := len(graph)
	d := make([][]int, n)
	for i := range d {
		d[i] = make([]int, n)
		for j := 0; j < n; j++ {
			d[i][j] = n
		}
	}

	for u, vs := range graph {
		for _, v := range vs {
			d[u][v] = 1
		}
	}

	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				d[i][j] = Min(d[i][j], d[i][k]+d[k][j])
			}
		}
	}

	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, 1<<n)
		for j := 0; j < 1<<n; j++ {
			f[i][j] = n * n
		}
	}
	for i := 0; i < n; i++ {
		f[i][1<<(n-1-i)] = 0
	}

	for mask := 0; mask < 1<<n; mask++ {
		for i := 0; i < n; i++ {
			if mask&(1<<(n-1-i)) == 0 {
				continue
			}
			for j := 0; j < n; j++ {
				if mask&(1<<(n-1-j)) != 0 {
					continue
				}
				f[j][mask^(1<<(n-1-j))] = Min(f[j][mask^(1<<(n-1-j))], f[i][mask]+d[i][j])
			}
		}
	}

	out := n * n
	for i := 0; i < n; i++ {
		out = Min(out, f[i][1<<n-1])
	}
	return out
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
