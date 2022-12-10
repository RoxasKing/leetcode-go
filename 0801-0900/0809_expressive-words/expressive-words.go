package main

// Difficulty:
// Medium

// Tags:
// Two Pointers

func expressiveWords(s string, words []string) int {
	m := len(s)
	o := 0
	for _, w := range words {
		ok := true
		n, i, j := len(w), 0, 0
		for i < m && j < n {
			if s[i] != w[j] {
				ok = false
				break
			}
			a, b := 1, 1
			for ; i+1 < m && s[i] == s[i+1]; i++ {
				a++
			}
			for ; j+1 < n && w[j] == w[j+1]; j++ {
				b++
			}
			if a < b || a == 2 && b == 1 {
				ok = false
				break
			}
			i++
			j++
		}
		if ok && i == m && j == n {
			o++
		}
	}
	return o
}
