package main

// Difficulty:
// Hard

// Tags:
// Prefix Sum
// Hash

func numSubmatrixSumTarget(matrix [][]int, target int) int {
	m, n := len(matrix), len(matrix[0])
	o := 0
	for r1 := 0; r1 < m; r1++ {
		f := make([]int, n)
		for r2 := r1; r2 < m; r2++ {
			for c := 0; c < n; c++ {
				f[c] += matrix[r2][c]
			}
			h := map[int]int{0: 1}
			v := 0
			for _, x := range f {
				v += x
				if cnt, ok := h[v-target]; ok {
					o += cnt
				}
				h[v]++
			}
		}
	}

	return o
}
