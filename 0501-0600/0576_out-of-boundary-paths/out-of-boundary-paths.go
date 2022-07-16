package main

// Difficulty:
// Medium

// Tags:
// Dynamic Programming

func findPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
	const mod int = 1e9 + 7
	dirs := []int{-1, 0, 1, 0, -1}
	f0 := [50][50]int{}
	f0[startRow][startColumn] = 1
	o := 0
	for k := 0; k < maxMove; k++ {
		f1 := [50][50]int{}
		for x0 := 0; x0 < m; x0++ {
			for y0 := 0; y0 < n; y0++ {
				if f0[x0][y0] == 0 {
					continue
				}
				for i := 0; i < 4; i++ {
					x, y := x0+dirs[i], y0+dirs[i+1]
					if x < 0 || m-1 < x || y < 0 || n-1 < y {
						o = (o + f0[x0][y0]) % mod
						continue
					}
					f1[x][y] = (f1[x][y] + f0[x0][y0]) % mod
				}
			}
		}
		f0 = f1
	}
	return o
}
