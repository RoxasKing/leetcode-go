package main

// Tags:
// Dynamic Programming

func maxPoints(points [][]int) int64 {
	m, n := len(points), len(points[0])
	f := make([]int, n)
	copy(f, points[0])

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			f[j] = Max(f[j-1]-1, f[j])
		}
		for j := n - 2; j >= 0; j-- {
			f[j] = Max(f[j+1]-1, f[j])
		}
		for j := 0; j < n; j++ {
			f[j] += points[i][j]
		}
	}

	var out int
	for i := range f {
		out = Max(out, f[i])
	}
	return int64(out)
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
