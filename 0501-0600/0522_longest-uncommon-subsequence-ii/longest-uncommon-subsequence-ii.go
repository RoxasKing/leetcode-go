package main

// Tags:
// Two Pointers

func findLUSlength(strs []string) int {
	n := len(strs)
	out := -1
	for i := 0; i < n; i++ {
		j := 0
		for ; j < n; j++ {
			if i == j {
				continue
			}
			if LCS(strs[i], strs[j]) {
				break
			}
		}
		if j == n {
			out = Max(out, len(strs[i]))
		}
	}
	return out
}

func LCS(a, b string) bool {
	if len(a) > len(b) {
		return false
	}
	i := 0
	for j := 0; i < len(a) && j < len(b); j++ {
		if a[i] == b[j] {
			i++
		}
	}
	return i == len(a)
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
