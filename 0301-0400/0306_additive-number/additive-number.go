package main

// Difficulty:
// Medium

// Tags:
// Backtracking

func isAdditiveNumber(num string) bool {
	var dfs func(int, int, int, int) bool
	dfs = func(i, a, b, c int) bool {
		if i == len(num) {
			return a != -1 && b != -1 && a+b == c
		}
		c = c*10 + int(num[i]-'0')
		return c != 0 && dfs(i+1, a, b, c) ||
			a == -1 && dfs(i+1, c, b, 0) ||
			a != -1 && b == -1 && dfs(i+1, a, c, 0) ||
			a != -1 && b != -1 && a+b == c && dfs(i+1, b, c, 0)
	}
	return dfs(0, -1, -1, 0)
}
