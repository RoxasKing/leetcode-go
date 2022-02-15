package main

// Difficulty:
// Easy

func luckyNumbers(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	minr := make([]int, m)
	maxc := make([]int, n)
	for i := 0; i < m; i++ {
		minr[i] = 1<<31 - 1
		for j := 0; j < n; j++ {
			minr[i] = Min(minr[i], matrix[i][j])
			maxc[j] = Max(maxc[j], matrix[i][j])
		}
	}
	out := []int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if minr[i] == maxc[j] && minr[i] == matrix[i][j] {
				out = append(out, minr[i])
			}
		}
	}
	return out
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
