package main

// Difficulty:
// Medium

// Tags:
// Dynamic Programming

func orderOfLargestPlusSign(n int, mines [][]int) int {
	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, n)
	}
	for _, x := range mines {
		f[x[0]][x[1]] = -1
	}
	for i := 0; i < n; i++ {
		c := 1
		for j := 0; j < n; j++ {
			if f[i][j] == -1 {
				c = 1
				continue
			}
			f[i][j] = c
			c++
		}
		c = 1
		for j := n - 1; j >= 0; j-- {
			if f[i][j] == -1 {
				c = 1
				continue
			}
			f[i][j] = min(f[i][j], c)
			c++
		}
	}
	o := 0
	for j := 0; j < n; j++ {
		c := 1
		for i := 0; i < n; i++ {
			if f[i][j] == -1 {
				c = 1
				continue
			}
			f[i][j] = min(f[i][j], c)
			c++
		}
		c = 1
		for i := n - 1; i >= 0; i-- {
			if f[i][j] == -1 {
				c = 1
				continue
			}
			f[i][j] = min(f[i][j], c)
			c++
			o = max(o, f[i][j])
		}
	}

	return o
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
