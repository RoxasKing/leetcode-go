package main

// Tags:
// Prefix Sum
// Hash

func numSubmatrixSumTarget(matrix [][]int, target int) int {
	m, n := len(matrix), len(matrix[0])
	pSum := make([][]int, m)
	for i := 0; i < m; i++ {
		pSum[i] = make([]int, n)
		for j := 0; j < n; j++ {
			pSum[i][j] = matrix[i][j]
			if i > 0 {
				pSum[i][j] += pSum[i-1][j]
			}
			if j > 0 {
				pSum[i][j] += pSum[i][j-1]
			}
			if i > 0 && j > 0 {
				pSum[i][j] -= pSum[i-1][j-1]
			}
		}
	}

	out := 0

	for r1 := 0; r1 < m; r1++ {
		for r2 := r1; r2 < m; r2++ {
			dict := map[int]int{0: 1}
			for c := 0; c < n; c++ {
				sum := pSum[r2][c]
				if r1 > 0 {
					sum -= pSum[r1-1][c]
				}
				if cnt, ok := dict[sum-target]; ok {
					out += cnt
				}
				dict[sum]++
			}
		}
	}

	return out
}
