package main

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}
	m, n := len(matrix), len(matrix[0])
	l, r, u, d := 0, n-1, 0, m-1
	out := make([]int, 0, m*n)
	f := 0
	for l <= r && u <= d {
		switch f {
		case 0:
			for j := l; j <= r; j++ {
				out = append(out, matrix[u][j])
			}
			u++
		case 1:
			for i := u; i <= d; i++ {
				out = append(out, matrix[i][r])
			}
			r--
		case 2:
			for j := r; j >= l; j-- {
				out = append(out, matrix[d][j])
			}
			d--
		case 3:
			for i := d; i >= u; i-- {
				out = append(out, matrix[i][l])
			}
			l++
		}
		f = (f + 1) % 4
	}
	return out
}
