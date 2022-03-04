package main

// Difficulty:
// Medium

// Tags:
// Dynamic Programming

func champagneTower(poured int, query_row int, query_glass int) float64 {
	f := make([][]float64, query_row+1)
	for i := range f {
		f[i] = make([]float64, query_row+1)
	}
	f[0][0] = float64(poured)
	for i := 0; i <= query_row; i++ {
		for j := 0; j <= i; j++ {
			if f[i][j] < 1.0 {
				continue
			}
			excess := (f[i][j] - 1.0) / 2.0
			f[i][j] = 1.0
			if i < query_row {
				f[i+1][j] += excess
				if j < query_row {
					f[i+1][j+1] += excess
				}
			}
			if i == query_row && j == query_glass {
				break
			}
		}
	}
	return f[query_row][query_glass]
}
