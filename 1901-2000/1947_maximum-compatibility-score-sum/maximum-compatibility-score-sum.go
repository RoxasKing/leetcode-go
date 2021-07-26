package main

// Tags:
// Backtracking

func maxCompatibilitySum(students [][]int, mentors [][]int) int {
	m, n := len(students), len(students[0])
	var out int
	backtrack(students, mentors, m, n, 0, &out)
	return out
}

func backtrack(students [][]int, mentors [][]int, m, n, i int, out *int) {
	if m == i {
		var cnt int
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if students[i][j] == mentors[i][j] {
					cnt++
				}
			}
		}
		*out = Max(*out, cnt)
		return
	}
	for j := i; j < m; j++ {
		mentors[i], mentors[j] = mentors[j], mentors[i]
		backtrack(students, mentors, m, n, i+1, out)
		mentors[i], mentors[j] = mentors[j], mentors[i]
	}
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
