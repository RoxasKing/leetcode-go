package main

// Tags:
// Dynamic Programming
func matrixScore(A [][]int) int {
	m, n := len(A), len(A[0])
	count := make([]int, n)
	for i := range A {
		for j := range A[i] {
			if A[i][j] == 1 {
				count[j]++
			}
		}
	}
	for i := range A {
		if A[i][0] == 1 {
			continue
		}
		for j := range A[i] {
			if A[i][j] == 0 {
				A[i][j] = 1
				count[j]++
			} else {
				A[i][j] = 0
				count[j]--
			}
		}
	}
	for j := range count {
		if count[j] >= (m+1)>>1 {
			continue
		}
		for i := range A {
			if A[i][j] == 0 {
				A[i][j] = 1
			} else {
				A[i][j] = 0
			}
		}
		count[j] = m - count[j]
	}
	out := 0
	for j := range count {
		out += count[j] * (1 << (n - 1 - j))
	}
	return out
}
