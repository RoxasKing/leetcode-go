package main

// Difficulty:
// Easy

func maxScore(s string) int {
	k := 0
	for _, c := range s {
		if c == '1' {
			k++
		}
	}
	o := 0
	for x, n, i := 0, len(s), 0; i < n-1; i++ {
		if s[i] == '1' {
			x++
		}
		o = max(o, i+1+k-x*2)
	}
	return o
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
