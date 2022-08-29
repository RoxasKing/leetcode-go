package main

// Difficulty:
// Hard

// Tags:
// Prefix Sum

func maxSumSubmatrix(matrix [][]int, k int) int {
	o, m, n := -1<<31, len(matrix), len(matrix[0])
	for x := 0; x < m; x++ {
		f := make([]int, n)
		for i := x; i < m; i++ {
			for j := 0; j < n; j++ {
				f[j] += matrix[i][j]
			}
			for l := 0; l < n; l++ {
				for r, sum := l, 0; r < n; r++ {
					if sum += f[r]; sum <= k && o < sum {
						o = sum
					}
				}
			}
		}
	}
	return o
}
