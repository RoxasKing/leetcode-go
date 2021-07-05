package main

// Tags:
// Backtracking

func isAdditiveNumber(num string) bool {
	return dfs(num, len(num), 0, -1, -1, 0)
}

func dfs(num string, n, i, p1, p2, cur int) bool {
	if i == n {
		return p1 != -1 && p2 != -1 && p1+p2 == cur
	}

	cur = cur*10 + int(num[i]-'0')

	if cur != 0 && dfs(num, n, i+1, p1, p2, cur) {
		return true
	}

	if p1 == -1 {
		return dfs(num, n, i+1, cur, p2, 0)
	}

	if p2 == -1 {
		return dfs(num, n, i+1, p1, cur, 0)
	}

	if p1+p2 == cur {
		return dfs(num, n, i+1, p2, cur, 0)
	}

	return false
}
