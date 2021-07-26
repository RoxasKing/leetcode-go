package main

// Tags:
// Dynamic Programming

func longestPalindrome(s string) string {
	n := len(s)
	f := make([][]bool, n)
	for i := range f {
		f[i] = make([]bool, n)
		f[i][i] = true
	}

	out := s[0:1]
	for size := 2; size <= n; size++ {
		for i := 0; i+size-1 < n; i++ {
			j := i + size - 1
			if s[i] != s[j] {
				continue
			}
			if j-i < 3 {
				f[i][j] = true
			} else {
				f[i][j] = f[i+1][j-1]
			}

			if f[i][j] && size > len(out) {
				out = s[i : j+1]
			}
		}
	}
	return out
}
