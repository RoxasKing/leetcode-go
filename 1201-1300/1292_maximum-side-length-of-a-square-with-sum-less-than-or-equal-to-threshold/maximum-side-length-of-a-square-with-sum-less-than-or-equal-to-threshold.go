package main

// Tags:
// Binary Search

func maxSideLength(mat [][]int, threshold int) int {
	m, n := len(mat), len(mat[0])
	pSum := make([][]int, m+1)
	for i := range pSum {
		pSum[i] = make([]int, n+1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			pSum[i+1][j+1] = mat[i][j] + pSum[i][j+1] + pSum[i+1][j] - pSum[i][j]
		}
	}

	out := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			l, r := 0, Min(m-i, n-j)
			for l < r {
				b := (l + r + 1) >> 1
				sum := pSum[i+b][j+b] - pSum[i+b][j] - pSum[i][j+b] + pSum[i][j]
				if sum <= threshold {
					l = b
				} else {
					r = b - 1
				}
			}
			out = Max(out, l)
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
