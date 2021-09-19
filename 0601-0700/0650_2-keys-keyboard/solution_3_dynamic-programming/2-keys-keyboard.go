package main

// Tags:
// Dynamic Programming

func minSteps(n int) int {
	f := make([]int, n+1)
	f[1] = 0
	for i := 2; i <= n; i++ {
		f[i] = n
		for k := 1; k <= i>>1; k++ {
			if i%k != 0 {
				continue
			}
			f[i] = Min(f[i], f[k]+i/k)
		}
	}
	return f[n]
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
