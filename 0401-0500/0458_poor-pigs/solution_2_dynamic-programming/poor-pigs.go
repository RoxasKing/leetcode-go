package main

// Difficulty:
// Hard

// Tags:
// Dynamic Programming
// Math

func poorPigs(buckets, minutesToDie, minutesToTest int) int {
	if buckets == 1 {
		return 0
	}
	c := make([][]int, buckets+1)
	for i := range c {
		c[i] = make([]int, buckets+1)
	}
	c[0][0] = 1
	t := minutesToTest / minutesToDie
	f := make([][]int, buckets)
	for i := range f {
		f[i] = make([]int, t+1)
	}
	for i := 0; i < buckets; i++ {
		f[i][0] = 1
	}
	for j := 0; j <= t; j++ {
		f[0][j] = 1
	}
	for i := 1; i < buckets; i++ {
		c[i][0], c[i][i] = 1, 1
		for j := 1; j < i; j++ {
			c[i][j] = c[i-1][j-1] + c[i-1][j]
		}
		for j := 1; j <= t; j++ {
			for k := 0; k <= i; k++ {
				f[i][j] += f[k][j-1] * c[i][i-k]
			}
		}
		if f[i][t] >= buckets {
			return i
		}
	}
	return 0
}
