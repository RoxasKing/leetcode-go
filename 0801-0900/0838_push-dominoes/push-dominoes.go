package main

// Difficulty:
// Medium

// Tags:
// Dynamic Programming

func pushDominoes(dominoes string) string {
	n := len(dominoes)
	l, r := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		l[i], r[i] = 1e5, 1e5
	}
	for i := 0; i < n; i++ {
		if dominoes[i] == 'R' {
			r[i] = 0
		} else if i > 0 && dominoes[i] == '.' {
			r[i] = min(r[i], r[i-1]+1)
		}
		if dominoes[n-1-i] == 'L' {
			l[n-1-i] = 0
		} else if i > 0 && dominoes[n-1-i] == '.' {
			l[n-1-i] = min(l[n-1-i], l[n-i]+1)
		}
	}
	a := make([]byte, n)
	for i := 0; i < n; i++ {
		a[i] = '.'
		if l[i] < r[i] {
			a[i] = 'L'
		}
		if l[i] > r[i] {
			a[i] = 'R'
		}
	}
	return string(a)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
