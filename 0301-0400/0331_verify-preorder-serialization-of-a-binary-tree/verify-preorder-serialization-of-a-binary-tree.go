package main

import (
	"strings"
)

// Tags:
// DFS
func isValidSerialization(preorder string) bool {
	strs := strings.Split(preorder, ",")
	i := 0
	dfs(strs, &i)
	return i == len(strs)-1
}

func dfs(strs []string, i *int) bool {
	if strs[*i] == "#" {
		return true
	}
	*i++
	if *i == len(strs) {
		return false
	}
	ok := dfs(strs, i)
	if !ok {
		return false
	}
	*i++
	if *i == len(strs) {
		return false
	}
	ok = dfs(strs, i)
	return ok
}
