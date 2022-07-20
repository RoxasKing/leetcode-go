package main

// Difficulty:
// Hard

// Tags:
// Prefix Sum
// Hash

func numSubmatrixSumTarget(matrix [][]int, target int) int {
	m, n := len(matrix), len(matrix[0])
	f := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		f[i] = make([]int, n+1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			f[i+1][j+1] = f[i+1][j] + f[i][j+1] - f[i][j] + matrix[i][j]
		}
	}
	o := 0
	for r1 := 1; r1 <= m; r1++ {
		for r2 := r1; r2 <= m; r2++ {
			h := map[int]int{0: 1}
			for c := 1; c <= n; c++ {
				v := f[r2][c] - f[r1-1][c]
				if cnt, ok := h[v-target]; ok {
					o += cnt
				}
				h[v]++
			}
		}
	}
	return o
}
