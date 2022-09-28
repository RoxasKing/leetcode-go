package main

// Difficulty:
// Hard

// Tags:
// Backtracking

func kSimilarity(s1 string, s2 string) int {
	a, b := []byte(s1), []byte(s2)
	n := len(a)
	o := n
	var backtrack func(i, k, c int)
	backtrack = func(i, k, c int) {
		if i == n || o <= k || c == 0 {
			if o > k {
				o = k
			}
			return
		}
		if a[i] == b[i] {
			backtrack(i+1, k, c)
			return
		}
		for j := i + 1; j < n; j++ {
			if a[j] == b[j] || a[i] != b[j] {
				continue
			}
			k++
			c--
			if a[j] == b[i] {
				c--
			}
			b[i], b[j] = b[j], b[i]
			backtrack(i+1, k, c)
			b[i], b[j] = b[j], b[i]
			c++
			if a[j] == b[i] {
				c++
			}
			k--
		}
	}
	backtrack(0, 0, n)
	return o
}
