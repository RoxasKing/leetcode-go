package main

// Tags:
// Dynamic Programming

func rearrangeSticks(n int, k int) int {
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, k+1)
	}
	f[0][0] = 1

	for i := 1; i <= n; i++ {
		for j := 1; j <= k && j <= i; j++ {
			f[i][j] = f[i-1][j]*(i-1) + f[i-1][j-1]
			f[i][j] %= 1e9 + 7
		}
	}

	return f[n][k]
}
