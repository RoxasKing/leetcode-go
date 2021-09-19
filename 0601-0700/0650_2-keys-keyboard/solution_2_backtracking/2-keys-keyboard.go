package main

// Tags:
// Backtracking

func minSteps(n int) int {
	out := n
	backtracking(n, 1, 0, 0, &out)
	return out
}

func backtracking(n, x, c, p int, out *int) {
	if c >= *out {
		return
	}

	if x == n {
		*out = Min(*out, c)
		return
	}

	if x<<1 <= n {
		backtracking(n, x<<1, c+2, x, out)
	}

	if p > 0 && x+p <= n {
		backtracking(n, x+p, c+1, p, out)
	}
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
