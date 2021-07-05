package main

import (
	"strings"
)

// Tags:
// Backtracking
func restoreIpAddresses(s string) []string {
	out := []string{}
	dfs(s, len(s), 0, []string{}, &out)
	return out
}

func dfs(s string, n, i int, cur []string, out *[]string) {
	if i == n {
		if len(cur) != 4 {
			return
		}
		*out = append(*out, strings.Join(cur, "."))
		return
	}

	if (4-len(cur))*3 < n-i {
		return
	}

	cur = append(cur, s[i:i+1])
	dfs(s, n, i+1, cur, out)
	cur = cur[:len(cur)-1]

	if s[i] == '0' || i+1 == n {
		return
	}

	cur = append(cur, s[i:i+2])
	dfs(s, n, i+2, cur, out)
	cur = cur[:len(cur)-1]

	if i+2 == n {
		return
	}

	if s[i] > '2' || s[i] == '2' && (s[i+1] > '5' || s[i+1] == '5' && s[i+2] > '5') {
		return
	}

	cur = append(cur, s[i:i+3])
	dfs(s, n, i+3, cur, out)
	// cur = cur[:len(cur)-1]
}
