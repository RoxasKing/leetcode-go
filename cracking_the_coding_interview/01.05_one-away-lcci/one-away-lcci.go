package main

// Difficulty:
// Medium

func oneEditAway(first string, second string) bool {
	m, n := len(first), len(second)
	if m-n > 1 || n-m > 1 {
		return false
	}
	if m < n {
		n = m
	}
	for i := 0; i < n; i++ {
		if first[i] != second[i] {
			return first[i+1:] == second[i+1:] || first[i+1:] == second[i:] || first[i:] == second[i+1:]
		}
	}
	return true
}
