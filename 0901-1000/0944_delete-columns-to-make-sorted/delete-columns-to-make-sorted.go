package main

// Difficulty:
// Easy

func minDeletionSize(strs []string) int {
	m, n := len(strs), len(strs[0])
	o := n
	for j := 0; j < n; j++ {
		i := 1
		for ; i < m && strs[i-1][j] <= strs[i][j]; i++ {
		}
		if i == m {
			o--
		}
	}
	return o
}
