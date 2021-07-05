package main

// Tags:
// Dynamic Programming
func matrixBlockSum(mat [][]int, K int) [][]int {
	m, n := len(mat), len(mat[0])
	sum := make([][]int, m)
	out := make([][]int, m)
	for i := range out {
		sum[i] = make([]int, n)
		out[i] = make([]int, n)
	}
	for i := range mat {
		for j := range mat[i] {
			sum[i][j] = mat[i][j]
			if i > 0 {
				sum[i][j] += sum[i-1][j]
			}
			if j > 0 {
				sum[i][j] += sum[i][j-1]
			}
			if i > 0 && j > 0 {
				sum[i][j] -= sum[i-1][j-1]
			}
		}
	}
	for i := range out {
		for j := range out[i] {
			a, b := Max(0, i-K), Max(0, j-K)
			c, d := Min(m-1, i+K), Min(n-1, j+K)
			out[i][j] = sum[c][d]
			if a > 0 {
				out[i][j] -= sum[a-1][d]
			}
			if b > 0 {
				out[i][j] -= sum[c][b-1]
			}
			if a > 0 && b > 0 {
				out[i][j] += sum[a-1][b-1]
			}
		}
	}
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
