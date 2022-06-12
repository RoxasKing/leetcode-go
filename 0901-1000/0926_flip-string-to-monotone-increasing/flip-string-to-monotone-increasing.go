package main

// Difficulty:
// Medium

// Tags:
// Dynamic Programming

func minFlipsMonoIncr(s string) int {
	n := len(s)
	cnt0 := 0
	for _, c := range s {
		if c == '0' {
			cnt0++
		}
	}
	o := min(cnt0, n-cnt0)
	cur0 := 0
	for i, c := range s {
		if c == '0' {
			cur0++
		}
		o = min(o, (i+1-cur0)+(cnt0-cur0))
	}
	return o
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
