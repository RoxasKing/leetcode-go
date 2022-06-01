package main

// Difficulty:
// Medium

// Tags:
// Sliding Window
// Bit Manipulation
// Hash

func hasAllCodes(s string, k int) bool {
	n := 1 << k
	f := make([]bool, n)
	c := 0
	for x, i := 0, 0; i < len(s); i++ {
		x = x<<1 + int(s[i]-'0')
		if i > k-1 {
			x &= n - 1
		}
		if i >= k-1 && !f[x] {
			f[x] = true
			c++
		}
	}
	return c == 1<<k
}
