package main

// Difficulty:
// Easy

func numSpecial(mat [][]int) int {
	m, n := len(mat), len(mat[0])
	rc, cc := make([]int, m), make([]int, n)
	for i, vs := range mat {
		for j, v := range vs {
			if v == 1 {
				rc[i]++
				cc[j]++
			}
		}
	}
	o := 0
	for i, vs := range mat {
		for j, v := range vs {
			if v == 1 && rc[i] == 1 && cc[j] == 1 {
				o++
			}
		}
	}
	return o
}
