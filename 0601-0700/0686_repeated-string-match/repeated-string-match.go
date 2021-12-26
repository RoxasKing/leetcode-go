package main

// Difficulty:
// Medium

// Tags:
// KMP

func repeatedStringMatch(a string, b string) int {
	m, n := len(a), len(b)
	pi := make([]int, n)
	for i, j := 1, 0; i < n; i++ {
		for ; j > 0 && b[i] != b[j]; j = pi[j-1] {
		}
		if b[i] == b[j] {
			j++
		}
		pi[i] = j
	}
	for i, j, cnt := 0, 0, 1; cnt <= 2 || m < n && m*cnt <= n*3; i = (i + 1) % m {
		for ; j > 0 && a[i] != b[j]; j = pi[j-1] {
		}
		if a[i] == b[j] {
			j++
		}
		if j == n {
			return cnt
		}
		if i == m-1 {
			cnt++
		}
	}
	return -1
}
