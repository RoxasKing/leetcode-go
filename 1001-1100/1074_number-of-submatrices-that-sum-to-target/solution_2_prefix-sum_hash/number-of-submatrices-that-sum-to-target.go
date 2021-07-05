package main

// Tags:
// Prefix Sum
// Hash

func numSubmatrixSumTarget(matrix [][]int, target int) int {
	m, n := len(matrix), len(matrix[0])
	out := 0

	for r1 := 0; r1 < m; r1++ {
		pSum := make([]int, n)
		for r2 := r1; r2 < m; r2++ {
			for c := 0; c < n; c++ {
				pSum[c] += matrix[r2][c]
			}
			dict := map[int]int{0: 1}
			cur := 0
			for _, sum := range pSum {
				cur += sum
				if cnt, ok := dict[cur-target]; ok {
					out += cnt
				}
				dict[cur]++
			}
		}
	}

	return out
}
