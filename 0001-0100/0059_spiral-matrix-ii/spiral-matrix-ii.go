package main

// Difficulty:
// Medium

// Tags:
// Simulation

func generateMatrix(n int) [][]int {
	o := make([][]int, n)
	for i := range o {
		o[i] = make([]int, n)
	}
	l, r, u, d := 0, n-1, 0, n-1
	dir := 0
	idx := 1
	for l <= r && u <= d {
		switch dir {
		case 0:
			for i := l; i <= r; i++ {
				o[u][i] = idx
				idx++
			}
			u++
		case 1:
			for i := u; i <= d; i++ {
				o[i][r] = idx
				idx++
			}
			r--
		case 2:
			for i := r; i >= l; i-- {
				o[d][i] = idx
				idx++
			}
			d--
		case 3:
			for i := d; i >= u; i-- {
				o[i][l] = idx
				idx++
			}
			l++
		}
		dir = (dir + 1) % 4
	}
	return o
}
