package main

func generateMatrix(n int) [][]int {
	out := make([][]int, n)
	for i := range out {
		out[i] = make([]int, n)
	}
	u, d, l, r := 0, n-1, 0, n-1
	num, move := 1, 0
	for u <= d && l <= r {
		switch move {
		case 0:
			for i := l; i <= r; i++ {
				out[u][i] = num
				num++
			}
			u++
		case 1:
			for i := u; i <= d; i++ {
				out[i][r] = num
				num++
			}
			r--
		case 2:
			for i := r; i >= l; i-- {
				out[d][i] = num
				num++
			}
			d--
		case 3:
			for i := d; i >= u; i-- {
				out[i][l] = num
				num++
			}
			l++
		}
		move = (move + 1) % 4
	}
	return out
}
