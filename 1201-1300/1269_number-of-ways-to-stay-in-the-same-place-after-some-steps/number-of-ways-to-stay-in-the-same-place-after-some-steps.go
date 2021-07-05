package main

// Tags:
// Dynamic Programming
func numWays(steps int, arrLen int) int {
	arrLen = Min(arrLen, steps+1)
	f := make([][]int, steps+1)
	for i := range f {
		f[i] = make([]int, arrLen)
	}
	f[0][0] = 1
	for i := 1; i <= steps; i++ {
		for j := 0; j < arrLen; j++ {
			f[i][j] = f[i-1][j]
			if j-1 >= 0 {
				f[i][j] += f[i-1][j-1]
				f[i][j] %= 1e9 + 7
			}
			if j+1 < arrLen {
				f[i][j] += f[i-1][j+1]
				f[i][j] %= 1e9 + 7
			}
		}
	}
	return f[steps][0]
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
