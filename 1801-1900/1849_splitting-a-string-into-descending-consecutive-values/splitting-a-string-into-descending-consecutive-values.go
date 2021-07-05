package main

import (
	"strconv"
)

// Tags:
// Backtracking
func splitString(s string) bool {
	return dfs(s, len(s), 0, []int{})
}

func dfs(s string, n, start int, cur []int) bool {
	if start == n {
		return len(cur) > 1
	}
	for i := start + 1; i <= n; i++ {
		num, _ := strconv.Atoi(s[start:i])
		if len(cur) > 0 && cur[len(cur)-1]-1 != num {
			continue
		}
		cur = append(cur, num)
		if dfs(s, n, i, cur) {
			return true
		}
		cur = cur[:len(cur)-1]
	}
	return false
}
