package main

import "strings"

// Difficulty:
// Medium

// Tags:
// Stack

func simplifyPath(path string) string {
	stk := []string{}
	n := len(path)
	for i, j := 0, 0; i < n; i = j {
		for ; i < n && path[i] == '/'; i++ {
		}
		for j = i; j < n && path[j] != '/'; j++ {
		}
		if i == j {
			continue
		}
		if path[i] == '.' && i+1 == j {
		} else if path[i] == '.' && i < n-1 && path[i+1] == '.' && i+2 == j {
			if len(stk) > 0 {
				stk = stk[:len(stk)-1]
			}
		} else {
			stk = append(stk, path[i:j])
		}
	}
	return "/" + strings.Join(stk, "/")
}
