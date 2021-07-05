package main

// Tags:
// Math

func waysToFillArray(queries [][]int) []int {
	f := make([][]int, 10000+13)
	for i := range f {
		f[i] = make([]int, 13+1)
	}
	f[0][0] = 1
	for i := 1; i < 10000+13; i++ {
		f[i][0] = 1
		for j := 1; j <= i && j <= 13; j++ {
			f[i][j] = f[i-1][j-1] + f[i-1][j]
			f[i][j] %= mod
		}
	}

	out := make([]int, len(queries))

	for i, q := range queries {
		n, k := q[0], q[1]
		out[i] = 1
		for j := 2; j*j <= k; j++ {
			if k%j > 0 {
				continue
			}
			r := 0
			for ; k%j == 0; k /= j {
				r++
			}
			out[i] *= f[r+n-1][r]
			out[i] %= mod
		}
		if k > 1 {
			out[i] *= n
			out[i] %= mod
		}
	}

	return out
}

var mod = int(1e9 + 7)
