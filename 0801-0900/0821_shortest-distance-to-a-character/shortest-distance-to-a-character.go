package main

// Difficulty:
// Easy

// Tags:
// Two Pointers

func shortestToChar(s string, c byte) []int {
	n := len(s)
	a := []int{}
	for i := range s {
		if s[i] == c {
			a = append(a, i)
		}
	}
	o := make([]int, n)
	for i, j := 0, 0; i < n; i++ {
		o[i] = n + 1
		for ; j < len(a) && a[j] < i; j++ {
		}
		if j > 0 {
			o[i] = Min(o[i], i-a[j-1])
		}
		if j < len(a) {
			o[i] = Min(o[i], a[j]-i)
		}
	}
	return o
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
