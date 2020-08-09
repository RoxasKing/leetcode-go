package leetcode

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
	var out []string
	var dfs func([]string, int)
	dfs = func(ip []string, l int) {
		if len(ip) == 3 {
			if valid(s[l:]) {
				ip = append(ip, s[l:])
				out = append(out, strings.Join(ip, "."))
			}
			return
		}
		maxRemain := 3 * (3 - len(ip))
		for r := l + 1; r <= l+3 && r < len(s); r++ {
			if r+maxRemain < len(s) {
				continue
			}
			if valid(s[l:r]) {
				ip = append(ip, s[l:r])
				dfs(ip, r)
				ip = ip[:len(ip)-1]
			}
		}
	}
	dfs(make([]string, 0, 4), 0)
	return out
}
