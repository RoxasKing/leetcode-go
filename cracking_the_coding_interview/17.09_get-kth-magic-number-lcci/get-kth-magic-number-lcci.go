package main

// Difficulty:
// Medium

// Tags:
// Dynamic Programming

func getKthMagicNumber(k int) int {
	f := make([]int, k)
	f[0] = 1
	for a, b, c, i := 0, 0, 0, 1; i < k; i++ {
		f[i] = min(f[a]*3, min(f[b]*5, f[c]*7))
		if f[a]*3 <= f[i] {
			a++
		}
		if f[b]*5 <= f[i] {
			b++
		}
		if f[c]*7 <= f[i] {
			c++
		}
	}
	return f[k-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
