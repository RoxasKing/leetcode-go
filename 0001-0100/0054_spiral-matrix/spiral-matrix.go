package main

func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	out := make([]int, 0, m*n)
	u, d, l, r := 0, m-1, 0, n-1
	forward := 0
	for u <= d && l <= r {
		switch forward {
		case 0:
			for i := l; i <= r; i++ {
				out = append(out, matrix[u][i])
			}
			u++
		case 1:
			for i := u; i <= d; i++ {
				out = append(out, matrix[i][r])
			}
			r--
		case 2:
			for i := r; i >= l; i-- {
				out = append(out, matrix[d][i])
			}
			d--
		case 3:
			for i := d; i >= u; i-- {
				out = append(out, matrix[i][l])
			}
			l++
		}
		forward = (forward + 1) % 4
	}
	return out
}
