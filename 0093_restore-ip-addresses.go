package main

import (
	"strings"
)

/*
  给定一个只包含数字的字符串，复原它并返回所有可能的 IP 地址格式。
*/

// DFS
func restoreIpAddresses(s string) []string {
	if len(s) < 4 || 12 < len(s) {
		return nil
	}
	valid := func(s string) bool {
		if len(s) > 1 && s[0] == '0' || len(s) == 3 &&
			(s[0] > '2' || s[0] == '2' && (s[1] > '5' || s[1] == '5' && s[2] > '5')) {
			return false
		}
		return true
	}
	cur := make([]string, 0, 4)
	var out []string
	var dfs func(int)
	dfs = func(L int) {
		if len(cur) == 3 {
			if valid(s[L:]) {
				cur = append(cur, s[L:])
				out = append(out, strings.Join(cur, "."))
				cur = cur[:len(cur)-1]
			}
			return
		}
		maxRemain := 3 * (3 - len(cur))
		for R := L + 1; R <= L+3 && R < len(s); R++ {
			if R+maxRemain < len(s) {
				continue
			}
			if valid(s[L:R]) {
				cur = append(cur, s[L:R])
				dfs(R)
				cur = cur[:len(cur)-1]
			}
		}
	}
	dfs(0)
	return out
}
