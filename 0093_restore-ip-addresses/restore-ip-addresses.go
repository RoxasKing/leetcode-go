package main

import (
	"strings"
)

/*
  给定一个只包含数字的字符串，复原它并返回所有可能的 IP 地址格式。
*/

// DFS + Backtracking
func restoreIpAddresses(s string) []string {
	out := []string{}
	cur := make([]string, 0, 3)
	dfs(s, 0, &cur, &out)
	return out
}

func dfs(s string, start int, cur, out *[]string) {
	if len(*cur) == 3 {
		if isValid(s[start:]) {
			*out = append(*out, strings.Join(*cur, ".")+"."+s[start:])
		}
		return
	}
	minR := len(s) - (3-len(*cur))*3
	for r := Max(start+1, minR); r <= start+3 && r < len(s); r++ {
		if isValid(s[start:r]) {
			*cur = append(*cur, s[start:r])
			dfs(s, r, cur, out)
			*cur = (*cur)[:len(*cur)-1]
		}
	}
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func isValid(s string) bool {
	n := len(s)
	if n > 1 && s[0] == '0' ||
		n == 3 && (s[0] > '2' || s[0] == '2' && (s[1] > '5' || s[1] == '5' && s[2] > '5')) {
		return false
	}
	return true
}
