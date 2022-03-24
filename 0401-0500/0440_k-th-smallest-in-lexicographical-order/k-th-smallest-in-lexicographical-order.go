package main

// Difficulty:
// Medium

// Tags:
// Trie

func findKthNumber(n int, k int) int {
	out := 1
	for k--; k > 0; {
		c := 0
		for l, r := out, out+1; l <= n; l, r = l*10, r*10 {
			c += Min(n+1, r) - l
		}
		if c > k {
			out *= 10
			k--
		} else {
			out++
			k -= c
		}
	}
	return out
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
