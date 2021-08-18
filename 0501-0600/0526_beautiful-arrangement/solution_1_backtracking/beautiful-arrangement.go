package main

// Tags:
// Backtracking

func countArrangement(n int) int {
	used := make([]bool, n)
	var out int
	backtrack(used, &out, n, 0)
	return out
}

func backtrack(used []bool, out *int, n, i int) {
	if n == i {
		(*out)++
		return
	}
	for j := 0; j < n; j++ {
		if used[j] {
			continue
		}
		if (j+1)%(i+1) == 0 || (i+1)%(j+1) == 0 {
			used[j] = true
			backtrack(used, out, n, i+1)
			used[j] = false
		}
	}
}
